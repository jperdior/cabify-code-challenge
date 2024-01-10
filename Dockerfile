# Use the official Golang base image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the Go application source code to the container
COPY . .

# Build the Go application
RUN go build -o main ./cmd/api

# Expose the port that the application will run on
EXPOSE 8080

# Run the Go application
CMD ["./main"]
