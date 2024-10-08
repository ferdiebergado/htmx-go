# Start with the official Go image
FROM golang:1.22.6-bookworm

ARG APP_ENV
ARG APP_PORT

# Install Air for hot reloading
RUN go install github.com/air-verse/air@latest

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum to cache dependencies
COPY go.mod ./
COPY go.sum ./

# Download dependencies
RUN go mod download

# Command to run Air
CMD ["air"]

# Expose the application port
EXPOSE ${APP_PORT}

# Mount your source files at /app when running the container
