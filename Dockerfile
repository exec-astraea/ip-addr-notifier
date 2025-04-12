FROM golang:1.24

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o ip-change-notifier .
CMD ["./ip-change-notifier"]
