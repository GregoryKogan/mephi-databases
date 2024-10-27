# Use the official Golang image as the base image
FROM golang:1.23.2-alpine

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY cmd/mephi-databases ./cmd/mephi-databases

# Build the Go app
RUN go build -o /mephi-databases ./cmd/mephi-databases

# Expose port 8888 to the outside world
EXPOSE 8888

# Command to run the executable
CMD ["/mephi-databases"]