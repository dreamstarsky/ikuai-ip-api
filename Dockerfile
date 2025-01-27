# Build stage
FROM golang:1.22-alpine AS build

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o ikuai-ip-api .

# Final stage
FROM alpine:latest

ENV GIN_MODE=release

WORKDIR /app

COPY --from=build /app/ikuai-ip-api /app/ikuai-ip-api

EXPOSE 8080

CMD ["./ikuai-ip-api"]
