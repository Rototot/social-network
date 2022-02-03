FROM golang:1.17-alpine as go_builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

#RUN apk add --no-cache git

RUN CGO_ENABLED=0 GOOS=linux \
    && go generate generate.go \
    && go generate ./pkg/... \
    && go build -o ./bin/social-network \
    && chmod +x ./bin/social-network


#### MAIN
FROM alpine:3.15

WORKDIR /app

COPY --from=go_builder /app/bin/social-network ./bin/

CMD ["/app/bin/social-network", "server"]

# 8000 = API
EXPOSE 8000

