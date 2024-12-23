# Use the official Golang image as a base
FROM golang:1.23.4

# Set the working directory
WORKDIR /app

# Copy application files
COPY . .

# Build the application
RUN go build -o app main.go

# Expose port 8080
EXPOSE 8080

# Run the application
CMD ["./app"]
