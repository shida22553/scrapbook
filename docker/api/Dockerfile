FROM golang:alpine3.13

RUN apk update
RUN apk add curl
RUN apk add build-base

RUN mkdir /usr/local/go/src/myapp
WORKDIR /usr/local/go/src/myapp
RUN go mod init shida22552/my-scrapbook
RUN go get -u github.com/cosmtrek/air
RUN go get github.com/gin-gonic/gin
RUN go get github.com/go-sql-driver/mysql
RUN go get github.com/jinzhu/gorm

COPY . /usr/local/go/src/myapp
