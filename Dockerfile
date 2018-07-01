FROM golang:1.10.3-alpine

LABEL maintainer "aimof <aimof.aimof@gmail.com>"

RUN apk update --no-cache

COPY . /go/src/github.com/aimof/catalpha/
WORKDIR /go/src/github.com/aimof/catalpha/cmd/catalpha/

RUN dep ensure && \
    go build

CMD ./catalpha
