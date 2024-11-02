# Start with a base image containing Go
FROM golang:1.23-alpine

# Set environment variables
ENV GO111MODULE=on

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod ./
COPY go.sum ./

# Download all Go dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the Go app
RUN go build -o main ./cmd/api/main.go

# Expose the port the app runs on
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
