# Start from the official Go image
FROM golang:1.16-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files
COPY go.mod go.sum ./

# Download and cache Go dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go application
RUN go build -o app

# Expose the port that the application listens on
EXPOSE 8080

# Run the Go application
CMD ["./app"]