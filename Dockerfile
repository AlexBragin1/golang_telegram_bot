FROM golang:1.21.3-alpine AS builder

COPY . /github.com/AlexBragin1/golang_telegram_bot/
WORKDIR /github.com/AlexBragin1/golang_telegram_bot/

RUN  go mod download &&  go build -o ./bin/bot cmd/bot/main.go

FROM alpine:latest

WORKDIR /root/
COPY --from=s@  github.com/AlexBragin1/golang_telegram_bot/bin/bot .
COPY --from=s@  github.com/AlexBragin1/golang_telegram_bot/bin/bot .
CMD ["*./bin"]
