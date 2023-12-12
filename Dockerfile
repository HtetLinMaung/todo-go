# Start from a Go image to build your application
FROM golang:alpine AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum and download dependencies (utilizes Docker cache)
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o todo ./cmd/todo

# Start a new stage for the final image
FROM alpine:latest  

# Set the Current Working Directory inside the container
WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /app/todo .

# Expose port 8080
EXPOSE 8080

# Command to run the executable
CMD ["./todo"]
