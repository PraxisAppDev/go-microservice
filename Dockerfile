FROM golang:1.22.2 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
EXPOSE 8080
RUN go build -o main .
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/main .
CMD ["./main"]
