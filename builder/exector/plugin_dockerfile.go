// Copyright (C) 2014-2018 Goodrain Co., Ltd.
// RAINBOND, Application Management Platform

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version. For any non-GPL usage of Rainbond,
// one or multiple Commercial Licenses authorized by Goodrain Co., Ltd.
// must be obtained first.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package exector

import (
	"fmt"
	"github.com/goodrain/rainbond-operator/util/constants"
	"os"
	"strings"

	"github.com/goodrain/rainbond/builder"
	"github.com/goodrain/rainbond/builder/sources"
	"github.com/goodrain/rainbond/util"

	"github.com/goodrain/rainbond/db"
	"github.com/goodrain/rainbond/event"

	"github.com/pquerna/ffjson/ffjson"

	"github.com/goodrain/rainbond/builder/model"

	"github.com/goodrain/rainbond/mq/api/grpc/pb"
	"github.com/sirupsen/logrus"
)

const (
	cloneTimeout    = 60
	buildingTimeout = 180
	formatSourceDir = "/cache/build/%s/source/%s"
)

func (e *exectorManager) pluginDockerfileBuild(task *pb.TaskMessage) {
	var tb model.BuildPluginTaskBody
	if err := ffjson.Unmarshal(task.TaskBody, &tb); err != nil {
		logrus.Errorf("unmarshal taskbody error, %v", err)
		return
	}
	eventID := tb.EventID
	logger := event.GetManager().GetLogger(eventID)
	logger.Info("从dockerfile构建插件任务开始执行", map[string]string{"step": "builder-exector", "status": "starting"})
	logrus.Info("start exec build plugin from image worker")
	defer event.GetManager().ReleaseLogger(logger)
	for retry := 0; retry < 2; retry++ {
		err := e.runD(&tb, logger)
		if err != nil {
			logrus.Errorf("exec plugin build from dockerfile error:%s", err.Error())
			logger.Info("dockerfile构建插件任务执行失败，开始重试", map[string]string{"step": "builder-exector", "status": "failure"})
		} else {
			return
		}
	}
	version, err := db.GetManager().TenantPluginBuildVersionDao().GetBuildVersionByDeployVersion(tb.PluginID, tb.VersionID, tb.DeployVersion)
	if err != nil {
		logrus.Errorf("get version error, %v", err)
		return
	}
	version.Status = "failure"
	if err := db.GetManager().TenantPluginBuildVersionDao().UpdateModel(version); err != nil {
		logrus.Errorf("update version error, %v", err)
	}
	MetricErrorTaskNum++
	logger.Error("dockerfile构建插件任务执行失败", map[string]string{"step": "callback", "status": "failure"})
}

func (e *exectorManager) runD(t *model.BuildPluginTaskBody, logger event.Logger) error {
	logger.Info("开始拉取代码", map[string]string{"step": "build-exector"})
	sourceDir := fmt.Sprintf(formatSourceDir, t.TenantID, t.VersionID)
	if t.Repo == "" {
		t.Repo = "master"
	}
	if !util.DirIsEmpty(sourceDir) {
		os.RemoveAll(sourceDir)
	}
	if err := util.CheckAndCreateDir(sourceDir); err != nil {
		return err
	}
	if _, _, err := sources.GitClone(sources.CodeSourceInfo{RepositoryURL: t.GitURL, Branch: t.Repo, User: t.GitUsername, Password: t.GitPassword}, sourceDir, logger, 4); err != nil {
		logger.Error("拉取代码失败", map[string]string{"step": "builder-exector", "status": "failure"})
		logrus.Errorf("[plugin]git clone code error %v", err)
		return err
	}
	if !checkDockerfile(sourceDir) {
		logger.Error("代码未检测到dockerfile，暂不支持构建，任务即将退出", map[string]string{"step": "builder-exector", "status": "failure"})
		logrus.Error("代码未检测到dockerfile")
		return fmt.Errorf("have no dockerfile")
	}

	logger.Info("代码检测为dockerfile，开始编译", map[string]string{"step": "build-exector"})
	mm := strings.Split(t.GitURL, "/")
	n1 := strings.Split(mm[len(mm)-1], ".")[0]
	buildImageName := fmt.Sprintf(builder.REGISTRYDOMAIN+"/plugin_%s_%s:%s", n1, t.PluginID, t.DeployVersion)
	logger.Info("start build image", map[string]string{"step": "builder-exector"})
	err := sources.ImageBuild("", sourceDir, util.GetenvDefault("RBD_NAMESPACE", constants.Namespace), t.PluginID, t.DeployVersion, logger, "plug-build", buildImageName, e.BuildKitImage, e.BuildKitArgs, e.BuildKitCache, e.KubeClient)
	if err != nil {
		logger.Error(fmt.Sprintf("build image %s failure,find log in rbd-chaos", buildImageName), map[string]string{"step": "builder-exector", "status": "failure"})
		logrus.Errorf("[plugin]build image error: %s", err.Error())
		return err
	}
	logger.Info("build image success, start to push image to local image registry", map[string]string{"step": "builder-exector"})
	logger.Info("push image success", map[string]string{"step": "build-exector"})
	version, err := db.GetManager().TenantPluginBuildVersionDao().GetBuildVersionByDeployVersion(t.PluginID, t.VersionID, t.DeployVersion)
	if err != nil {
		logrus.Errorf("get version error, %v", err)
		return err
	}
	version.BuildLocalImage = buildImageName
	version.Status = "complete"
	if err := db.GetManager().TenantPluginBuildVersionDao().UpdateModel(version); err != nil {
		logrus.Errorf("update version error, %v", err)
		return err
	}
	logger.Info("build plugin version by dockerfile success", map[string]string{"step": "last", "status": "success"})
	return nil
}

func checkDockerfile(sourceDir string) bool {
	if _, err := os.Stat(fmt.Sprintf("%s/Dockerfile", sourceDir)); os.IsNotExist(err) {
		return false
	}
	return true
}
