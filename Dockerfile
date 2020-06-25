FROM golang:alpine3.12

RUN apk add git
RUN go get golang.org/x/tools/cmd/godoc