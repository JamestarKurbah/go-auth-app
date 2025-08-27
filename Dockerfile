# Start from the official Golang image
FROM golang:1.23-alpine
# Set working directory
WORKDIR /app

# Install git (for go mod)
RUN apk add --no-cache git

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the Go app
RUN go build -o go-auth-app ./src/main.go

# Expose port 8080
EXPOSE 8080

# Start the app
CMD ["./go-auth-app"]
