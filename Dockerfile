# Use an official Go runtime as the base image
FROM golang:1.16-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the Go application code into the container
COPY . .

# Build the Go application
RUN go build -o main .

# Expose a port (if required)
EXPOSE 8080

# Run the Go application
CMD ["./main"]