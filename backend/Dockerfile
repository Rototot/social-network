FROM golang:1.17-alpine as go_builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download


COPY generate.go  ./
RUN CGO_ENABLED=0 GOOS=linux \
    && go generate generate.go

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux \
    && go generate generate.go \
    && go build -o ./bin/social-network \
    && chmod +x ./bin/social-network


#### MAIN
FROM alpine:3.15

RUN apk --no-cache add \
    curl

HEALTHCHECK --interval=30s --timeout=30s \
    CMD curl -f http://localhost:8000/ || exit 1

WORKDIR /app

COPY --from=go_builder /app/bin/social-network ./bin/
COPY --from=go_builder /app/db/ ./db/

CMD ["/app/bin/social-network", "server"]

# 8000 = API
EXPOSE 8000

