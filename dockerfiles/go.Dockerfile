FROM golang:1.7-alpine

RUN apk add --update git

COPY ./src/bin /go/bin

WORKDIR /go/bin
CMD ./app
