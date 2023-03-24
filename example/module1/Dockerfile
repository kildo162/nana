# Base image for Go build environment
FROM golang:1.19.1-alpine3.16 AS build-env

# Set environment variables
ENV GO111MODULE=on \
    WORKDIR=/go/src/app

# Set working directory
WORKDIR $WORKDIR

# Copy application source code
COPY . $WORKDIR

# Download Go module dependencies
RUN go mod download


# Build the application
RUN go build -o /go/bin/app

# Final image
FROM alpine:3.16

COPY --from=build-env /go/bin/app /usr/local/bin/app

# Start the application
CMD /usr/local/bin/app
