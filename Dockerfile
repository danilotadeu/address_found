FROM golang:1.17.6-alpine
# Add a work directory
WORKDIR /app
# Cache and install dependencies
COPY go.mod go.sum ./
RUN go mod download
# Copy app files
COPY . .
# Build app
RUN go build
# Expose port
EXPOSE 3000
# Start app
CMD ./r-customer-code-information