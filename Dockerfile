# Build stage
FROM golang:alpine3.18 AS builder

# Set environment variables for Go
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
ENV GOPROXY=https://proxy.golang.org,direct

# Set the working directory inside the container
WORKDIR /go/src/app

# Copy go.mod and go.sum first to leverage Docker cache
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the Go binary with optimizations for a smaller binary size
RUN go build -ldflags="-s -w" -o ./bin/nutrition ./main.go

# Final stage (minimal image for running the app)
FROM alpine:3.17

# Install required packages (ca-certificates for HTTPS)
RUN apk --no-cache add ca-certificates

# Set the working directory for the final container
WORKDIR /usr/bin

# Copy the built binary and environment file from the builder stage
COPY --from=builder /go/src/app/bin/nutrition /go/bin/
COPY --from=builder /go/src/app/.env ./

# Expose the port for the application
EXPOSE 1010

# Command to run the binary when the container starts
CMD ["/go/bin/nutrition", "--port", "1010"]
