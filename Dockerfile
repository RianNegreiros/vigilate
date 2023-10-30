# Use the official Go image as the base image.
FROM golang:1.21 AS build

# Set the working directory inside the container.
WORKDIR /app

# Copy the Go module files and download dependencies.
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application source code.
COPY . .

# Build the Go application with compiler optimizations.
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /app/vigilate ./cmd

# Create a minimal base image to reduce the image size.
FROM alpine:3.14

# Set the working directory inside the container.
WORKDIR /app

# Copy the binary from the build stage.
COPY --from=build /app/vigilate .

# Copy the migration files into the container.
COPY internal/database/migrations /app/internal/database/migrations

# Expose the port your application will listen on.
EXPOSE 8080

# Start your Go application.
CMD ["./vigilate"]
