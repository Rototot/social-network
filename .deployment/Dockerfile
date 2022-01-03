FROM golang:1.17-alpine as go_builder

WORKDIR /app

COPY backend/go.mod backend/go.sum ./
COPY backend/go.sum ./
RUN go mod download

COPY backend/ ./
RUN CGO_ENABLED=0 GOOS=linux \
    && go build -o ./bin/social-network \
    && chmod +x ./bin/social-network

FROM node:16-alpine as frontend_builder

WORKDIR /app

COPY frontend/package.json frontend/package-lock.json ./
RUN npm ci

COPY frontend/ ./
RUN npm run build

#### MAIN
FROM alpine:3.15

RUN apk --update add --no-cache \
    nginx \
    supervisor

WORKDIR /app

RUN mkdir -p ./logs/supervisor

COPY .deployment/supervisord/conf.d/ /etc/supervisor.d/
COPY .deployment/nginx/conf.d/ /etc/nginx/conf.d/

COPY --from=go_builder /app/bin/social-network ./bin/
COPY --from=frontend_builder /app/build/ ./public/

CMD ["/usr/bin/supervisord", "--nodaemon"]

# 8000 = API
# 80 = Nginx
EXPOSE 8000 80

