FROM golang:1.12.6

ENV GOPATH $GOPATH:/go
ENV PATH $PATH:$GOPATH/bin
WORKDIR /go/src/app

RUN apt-get update
RUN go get "github.com/go-sql-driver/mysql"
RUN go get "github.com/beego/bee"
RUN go get "github.com/astaxie/beego"