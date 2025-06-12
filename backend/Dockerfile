# Use the official Golang image as the builder
FROM golang:1.22 AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Install CA certificates
RUN apt-get update && apt-get install -y ca-certificates && update-ca-certificates

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app with static linking
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Use scratch for a minimal image
FROM scratch

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/main .

# Copy CA certificates to the scratch image
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Expose port 3000 to the outside world
EXPOSE 3000

# Command to run the executable
CMD ["/main"]
