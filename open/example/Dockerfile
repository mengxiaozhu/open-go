FROM golang:latest

COPY ./ $GOPATH/src/example

WORKDIR $GOPATH/src/example

# TimeZone
RUN cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime

# Build
RUN go get -v && go build simple.go

# Chmod
RUN chmod -R 777 *

EXPOSE 4000

CMD ["./simple"]