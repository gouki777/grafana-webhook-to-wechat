FROM golang:1.16.3-alpine3.12 as small
RUN set -ex; \
    ln -snf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && echo "Asia/Shanghai" > /etc/timezone  && \
    sed -i 's/dl-cdn.alpinelinux.org/mirrors.cloud.tencent.com/g' /etc/apk/repositories && \ 
    apk add --no-cache git && \
    git clone https://github.com/gouki777/grafana-webhook-to-wechat.git && \
    cd grafana-webhook-to-wechat && \
    export GO111MODULE=on && \
    go env -w GOPROXY=https://goproxy.io,direct && \
    go build -v -o /wechat main.go 
FROM alpine:3.12.7
COPY --from=small /wechat /
EXPOSE 8081
CMD ["/wechat"]