FROM golang:1.25.5-alpine3.23 AS builder
WORKDIR /app

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o myapp

FROM alpine:latest
COPY --from=builder /app/myapp .
CMD ["./myapp"]
EXPOSE 8080