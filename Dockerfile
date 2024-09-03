# Dockerfile
FROM golang:1.22.6


WORKDIR /app

COPY . .

RUN go mod tidy
RUN go build -o product-app ./cmd/main.go

CMD ["./product-app"]