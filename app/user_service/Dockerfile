# Build stage
FROM golang:1.20-alpine AS builder
WORKDIR /app

# Copy go.mod and go.sum
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the service
RUN go build -o user-service .

# Final stage
FROM alpine:3.16
WORKDIR /app
COPY --from=builder /app/user-service /app/

# Expose the port the service listens to
EXPOSE 50051

# Run the service
CMD ["./user-service"]
