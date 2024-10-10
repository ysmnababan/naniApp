# Build stage
FROM golang:1.22.2-alpine AS builder
WORKDIR /app

# Copy go.mod and go.sum
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the API Gateway
RUN go build -o api-gateway .

# Final stage
FROM alpine:3.16
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/api-gateway .

# Copy the .env file to the final image
COPY .env /app/.env

# Expose the port the API Gateway listens to
EXPOSE 8080

# Run the API Gateway
CMD ["./api-gateway"]
