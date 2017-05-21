FROM golang:latest

ADD . /go/src/api-skeleton

RUN glide install

ENTRYPOINT /go/bin/api-skeleton

EXPOSE 3000


