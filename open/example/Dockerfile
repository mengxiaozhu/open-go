FROM golang:latest

COPY ./ $GOPATH/src/github.com/mengxiaozhu/open-go/open/example

WORKDIR $GOPATH/src/github.com/mengxiaozhu/open-go/open/example

# TimeZone
RUN cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime

# Build
RUN go get -v && go build

# Chmod
RUN chmod -R 777 *

EXPOSE 4000

CMD ["./simple"]