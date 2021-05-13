FROM golang:alpine AS builder

RUN mkdir /app
ADD . /app
WORKDIR /app

RUN go build -o server .

ENTRYPOINT "./server"
EXPOSE 8080
