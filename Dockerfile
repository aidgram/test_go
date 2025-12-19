FROM golang:1.25.5-alpine3.23 AS builder
WORKDIR /app
COPY . .
RUN go build -o myapp .

FROM alpine:latest
COPY --from=builder /app/myapp .
CMD ["./myapp"]
EXPOSE 8080