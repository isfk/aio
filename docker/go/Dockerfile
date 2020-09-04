FROM golang:1.14-alpine

RUN echo "http://mirrors.aliyun.com/alpine/v3.10/main/" > /etc/apk/repositories

RUN apk update && apk add --no-cache git openssh-client expect wget ca-certificates bash make tzdata gcc g++ && \
    ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    echo "Asia/Shanghai" > /etc/timezone

RUN adduser -D -g '' appuser

WORKDIR /go/src/myapp/

VOLUME /go/src/myapp/

ENV GO111MODULE=on