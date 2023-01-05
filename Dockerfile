# Alpine is chosen for its small footprint
# compared to Ubuntu
FROM golang:1.18-alpine AS builder
LABEL maintainer="Alif Iftitah<alifipa5@gmail.com>"

ENV APP_NAME=todo_app
ENV GO111MODULE=on
ENV GOPRIVATE=github.com/ifty123
ENV TZ=Asia/Jakarta
ENV GIT_TERMINAL_PROMPT=0
ENV CGO_ENABLED=0

RUN apk update && apk upgrade
RUN apk add --no-cache --virtual .build-deps --no-mysql -q \
    bash \
    curl \
    busybox-extras \
    make \
    git \
    tzdata && \
    cp /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone
RUN apk update && apk add --no-cache coreutils

WORKDIR /app

RUN mkdir -p /app/todo
COPY . /app/todo
WORKDIR /app/todo

RUN ls -ls

RUN go mod tidy -compat=1.18

RUN go build

EXPOSE 3030

CMD "./todo_app"
