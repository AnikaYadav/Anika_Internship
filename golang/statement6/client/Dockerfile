FROM golang:1.25 AS builder
WORKDIR /app
COPY .. .
RUN go build -o client ./client

FROM debian:bookworm-slim
WORKDIR /app
COPY --from=builder /app/client .
CMD ["sleep", "infinity"]