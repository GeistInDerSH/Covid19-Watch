FROM golang:1.14

WORKDIR /go/src/app
COPY . .

RUN go build main.go
ENTRYPOINT /go/src/app/main

EXPOSE 80
