# Build stage
FROM golang:1.22.2-alpine AS builder
WORKDIR /app

# Copy go.mod and go.sum
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the service
RUN go build -o message-service .

# Final stage
FROM alpine:3.16
WORKDIR /app
COPY --from=builder /app/message-service /app/

# Copy the .env file to the final image
COPY .env /app/.env

# Expose the port the service listens to
EXPOSE 50052

# Run the service
# CMD ["./message-service"]
CMD ["sh", "-c", "ls -la && ./message-service"]
