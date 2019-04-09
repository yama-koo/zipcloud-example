FROM golang:1.12.2-alpine3.9

RUN apk add --no-cache ca-certificates \
    git \
    curl \
    gcc
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

WORKDIR /go/src/github.com/yama-koo/zipcloud-example/

COPY . .

ENV PATH=$PATH:/go/bin
ENV CGO_ENABLED=0

RUN dep ensure

