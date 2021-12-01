FROM golang:1.17 AS builder

COPY . /github.com/1makarov/go-logger-rabbitmq-example/
WORKDIR /github.com/1makarov/go-logger-rabbitmq-example/

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o ./.bin/app ./cmd/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /github.com/1makarov/go-logger-rabbitmq-example/.bin/app .

CMD ["./app"]