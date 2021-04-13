FROM golang:1.16.3

RUN mkdir /myapp
WORKDIR /myapp

RUN go mod init shida22552/my-scrapbook
RUN go get github.com/gin-gonic/gin
COPY . /myapp

CMD go run main.go