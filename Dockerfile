FROM golang:1.18-rc-alpine as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o ./bin/social-network \
    && chmod +x ./bin/social-network


FROM alpine:3.15

WORKDIR /app

COPY --from=builder /app/bin/social-network ./

CMD ["./social-network", "server"]

EXPOSE 8000