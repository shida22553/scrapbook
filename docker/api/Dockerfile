FROM golang:1.16.3

RUN mkdir /myapp
WORKDIR /myapp

RUN go mod init shida22552/my-scrapbook
RUN go get github.com/gin-gonic/gin
RUN go get github.com/go-sql-driver/mysql
RUN go get github.com/jinzhu/gorm
RUN go get github.com/cosmtrek/air

COPY . /myapp
