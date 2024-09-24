# Use an official Go runtime as a parent image
FROM golang:1.22 AS builder

# Set the working directory inside the container
WORKDIR /app

# Set environment variables for cross-compilation
# Adjust these to match the target architecture, e.g., 'linux/arm64' or 'linux/amd64'
ENV GOOS=linux
ENV GOARCH=arm64
ENV CGO_ENABLED=0

# Copy the go.mod and go.sum to leverage Docker cache. Security: avoid global pattern / recursive copy
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy the current directory contents into the container at /app. Security: avoid global pattern / recursive copy
COPY internal/ ./internal/
COPY config/ ./config/

# Build the Go app with cross-compilation settings. Using -ldflags="-s -w" removes debugging information, reducing binary size.
RUN go build -ldflags="-s -w" -o gomyloader ./cmd/...

# Use a smaller base image to run the compiled binary
FROM alpine:3.20  
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the binary and config file from the builder stage to the production image
COPY --from=builder /app/gomyloader .
COPY --from=builder /app/config/config.yaml ./config/config.yaml

EXPOSE 2112

# Run the web service on container startup.
CMD ["./gomyloader"]
