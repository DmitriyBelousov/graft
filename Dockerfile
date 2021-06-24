ARG GOVER=1.16.3-alpine
FROM golang:$GOVER as builder

ENV GO111MODULE=auto \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

COPY . /go/src/github.com/DmitriyBelousov/kek
WORKDIR /go/src/github.com/DmitriyBelousov/kek

RUN go build ~/go/src/github.com/DmitriyBelousov/kek/main.go
#COPY . /go/src/github.com/DmitriyBelousov/kek
RUN  ./main


#FROM golang:alpine AS builder
#ENV GO111MODULE=on
#    CGO_ENABLED=0
#    GOOS=linux
#    GOARCH=amd64
#WORKDIR /build
#COPY . .
#RUN go build -o main .
#WORKDIR /dist
#RUN cp /build/main .
#FROM scratch
#ENTRYPOINT ["/main"]
