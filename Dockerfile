# Stage 1: Build Golang Application
FROM golang:alpine
RUN apk update && apk add --no-cache git
# Timezone
ENV TZ Asia/Jakarta
# Set the working directory inside the container
WORKDIR /app
# Copy the source code to the container
COPY . .
# Copy go.mod and go.sum files
COPY go.mod ./
# Download Go dependencies
RUN go mod tidy && go mod download
# Build the application
RUN go build -o main ./server/cmd/
# Expose the necessary ports
EXPOSE 8080
# Run the application
CMD ["./main"]