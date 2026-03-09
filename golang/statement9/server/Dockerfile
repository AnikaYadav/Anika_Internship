# Build stage
FROM golang:1.25 AS builder
WORKDIR /app
COPY .. .
RUN go build -o server ./server

# Runtime stage
FROM gcr.io/distroless/base-debian12
WORKDIR /app
COPY --from=builder /app/server .
EXPOSE 50051
CMD ["./server"]