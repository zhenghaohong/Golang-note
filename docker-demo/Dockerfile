FROM golang

MAINTAINER Razil "zhengteshuai"

WORKDIR $GOPATH/src/note/docker-demo

ADD . $GOPATH/src/note/docker-demo

RUN go build main.go

EXPOSE 8888

ENTRYPOINT ["./main"]
