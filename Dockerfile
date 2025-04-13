FROM golang:1.24.2-alpine3.21 AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o ip-addr-notifier .

FROM alpine:3.21.3

WORKDIR /app
COPY --from=builder /app/ip-addr-notifier .

CMD ["./ip-addr-notifier"]
