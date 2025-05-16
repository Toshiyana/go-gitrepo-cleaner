FROM golang:1.24-alpine

WORKDIR /app

# Install git and other dependencies
RUN apk add --no-cache git

# Copy go.mod and go.sum files
COPY go.mod go.sum* ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN go build -o /go/bin/githubcli ./cmd/githubcli

# Set the entrypoint
ENTRYPOINT ["/go/bin/githubcli"]
CMD ["--help"]
