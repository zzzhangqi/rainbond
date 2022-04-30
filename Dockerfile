FROM golang:1.15 as build
WORKDIR /work
ARG APPNAME
ENV GOARCH amd64
ENV GOOS linux
ENV CGO_ENABLED 0
ENV GOPROXY https://goproxy.cn
ENV GO111MODULE on
ADD . .
RUN go mod vendor
RUN go mod tidy
RUN go build -trimpath -ldflags "-w -s -X github.com/goodrain/rainbond/cmd.version=v1" -o app ./cmd/$APPNAME

FROM alpine:3.12
COPY --from=build /work/app /
ENTRYPOINT ["/app"]