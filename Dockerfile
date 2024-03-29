FROM golang:latest
RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN go build main.go
EXPOSE 8090
CMD ["/app/main"]
