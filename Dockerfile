FROM golang:latest
MAINTAINER "QZQ<mad7802004@foxmail.com>"

ENV GO111MODULE on
WORKDIR $GOPATH/src/github.com/qzq1111/mangmang
COPY . $GOPATH/src/github.com/qzq1111/mangmang

RUN go mod download
RUN go mod vendor
RUN go build .

EXPOSE 80
ENTRYPOINT ["./mangmang"]