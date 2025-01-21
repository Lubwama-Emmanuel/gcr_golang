# Build Stage
FROM golang:1.23 as build

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum first (for caching dependencies)
COPY go.mod ./

# Download dependencies
# RUN go mod download

# Copy the rest of the source code
COPY . .

# Disable CGO for a static binary
ENV CGO_ENABLED=0

# Build the Go binary
RUN go build -o app .

# Run Stage
FROM alpine:3.18

# Set the working directory inside the container
WORKDIR /app

# Copy the compiled binary from the build stage
COPY --from=build /app/app .

# Expose the port your app runs on (optional but recommended)
EXPOSE 8080

# Command to run the app
CMD ["./app"]
