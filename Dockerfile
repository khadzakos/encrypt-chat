# Use the official Golang image as a builder
FROM golang:1.22-alpine AS builder

# Set the current working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files to the container
COPY go.mod go.sum ./

# Download all Go modules
RUN go mod download

# Source .env file
ENV ENV_FILE .env

# Copy the rest of the project files to the container
COPY . .

# Build the Go app
RUN go build -o main .

# Use a minimal image for the final build
FROM alpine:latest

# Set the current working directory inside the container
WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /app/main .

# Expose the port the application runs on
EXPOSE 8080

# Command to run the Go app
CMD ["./main"]