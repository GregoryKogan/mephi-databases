# Use the official Golang image as the base image
FROM golang:1.23.2-alpine

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY cmd ./cmd
COPY internal ./internal

# Copy the config file
COPY config.yml .

# Build the Go app
RUN go build -o /mephi-databases ./cmd/mephi-databases

# Command to run the executable
CMD ["/mephi-databases"]