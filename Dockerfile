# Use golang alpine image as the builder stage
FROM golang:1.22.4-alpine3.20 AS builder

# Install git and other necessary packages
RUN apk update && apk add --no-cache git

# Set the Current Working Directory inside the container
WORKDIR /src

# Copy everything from the current directory to the PWD(Present Working Directory) inside the container
COPY . .

# Fetch dependencies using go mod if your project uses Go modules
RUN go mod download

# Build the Go app
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o /usr/local/bin/decrypt

# Use Ubuntu as the final image
FROM ubuntu:latest

# Install Common Dependencies
RUN apt-get update && \
    apt install -y \
    ca-certificates \
    curl \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*

# Install etcdctl
RUN curl -L https://github.com/etcd-io/etcd/releases/download/v3.5.9/etcd-v3.5.9-linux-amd64.tar.gz \
    | tar xz --strip-components=1 -C /usr/local/bin etcd-v3.5.9-linux-amd64/etcdctl

# Set the Current Working Directory inside the container
WORKDIR /root/

# Copy our static executable
COPY --from=builder /usr/local/bin/decrypt /usr/local/bin/decrypt
RUN chmod +x /usr/local/bin/decrypt

# Set default command
CMD ["/bin/bash"]
