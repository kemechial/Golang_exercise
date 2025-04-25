# Use an official Golang runtime as a parent image
FROM golang:1.22-alpine

# Install bash
RUN apk update && apk add --no-cache bash

# Set the working directory inside the container
WORKDIR /app

# Keep the container running indefinitely
CMD ["tail", "-f", "/dev/null"]
