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
RUN go get google.golang.org/api/option
RUN go get github.com/gin-contrib/cors
RUN go get firebase.google.com/go

COPY . /usr/local/go/src/myapp
