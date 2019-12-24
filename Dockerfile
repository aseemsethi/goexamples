# Dockerfile References: https://docs.docker.com/engine/reference/builder/

# Start from the latest golang base image
FROM golang:latest

RUN echo 'Making Docker Image for JWT Server'

# Add Maintainer Info
LABEL maintainer="Aseem Sethi <aseemsethi70@gmail.com>"

# Set the Current Working Directory inside the container
WORKDIR /go/src

# Copy go mod and sum files
ADD jwt /go/src/jwt
ADD github.com /go/src/github.com

# Set the Current Working Directory inside the container
WORKDIR /go/src/jwt

# Build the Go pp
RUN go build

# Expose port 8000 to the outside world
EXPOSE 8000

# Command to run the executable
CMD ["./jwt"]
