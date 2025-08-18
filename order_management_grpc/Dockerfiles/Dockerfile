# Build stage
FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY .. .
RUN go mod tidy && go build -o server ./cmd/server/main.go

# Final stage
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/server ./server
EXPOSE 50051
CMD ["./server"]
