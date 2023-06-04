# Stage 1: Build Golang Application
FROM golang:alpine3.18 as builder
RUN apk update && apk add --no-cache git
WORKDIR /app
COPY . .

# Stage 2: Build small images
FROM alpine3.18
# Set the working directory inside the container
WORKDIR /app
# Copy the executable from the builder stage
COPY --from=builder /app/main /app
# Copy go.mod and go.sum files
COPY go.mod ./
# Download Go dependencies
RUN go mod download
# Copy the source code to the container
COPY . .
# Build the application
RUN go build -o ./server/cmd/
# Expose the necessary ports
EXPOSE 8080
# Run the application
CMD ["./main"]