FROM golang:1.16.5-alpine
# Add a work directory
WORKDIR /app
# Cache and install dependencies
COPY go.mod go.sum ./
RUN go mod download
# Copy app files
COPY . .
# Expose port
EXPOSE 3000
# Start app
CMD go run main.go