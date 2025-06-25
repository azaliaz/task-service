# syntax=docker/dockerfile:1.4
FROM golang:alpine AS builder
ARG VERSION

WORKDIR /src

# Настройка кешей для модулей и сборки
RUN go env -w GOCACHE=/cache/go-cache
RUN go env -w GOMODCACHE=/cache/gomod-cache

COPY ./go.* ./

RUN --mount=type=cache,target=/cache/gomod-cache \
    go mod download

COPY . .

RUN --mount=type=cache,target=/cache/gomod-cache --mount=type=cache,target=/cache/go-cache \
    CGO_ENABLED=0 GOOS=linux go build -mod=vendor -a -installsuffix cgo -o app -ldflags "-X 'main.version=${VERSION}'" ./cmd/task-service

# Финальный минимальный образ
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY --from=builder /src/app .

EXPOSE 8080

ENTRYPOINT ["./app"]
CMD ["run"]
