FROM golang:1.18-rc-alpine as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download
RUN go build -o ./app/bin/