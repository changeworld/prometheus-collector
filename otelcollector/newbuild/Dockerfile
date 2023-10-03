# Use an official Golang runtime as the base image
FROM golang:1.16 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go source code into the container
COPY . .

# Build the Go executable
RUN go build -o my-container-app

# Use a minimal base image for the final container
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the built executable from the builder image
COPY --from=builder /app/my-container-app /app/my-container-app

# Expose the port on which the application listens
EXPOSE 8080

# Run the Go executable
CMD ["./poc_app"]
