FROM golang:1.22-alpine

WORKDIR /app

COPY . .

ENV GIN_MODE=release

RUN go mod tidy

RUN go build -o ikuai-ip-api .

CMD ["./ikuai-ip-api"]

EXPOSE 8080
