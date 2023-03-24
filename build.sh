#!/bin/bash

# Install the nana binary
go mod tidy

# Build the project golang binary
go build -o bin/nana

# Go install the binary
sudo cp bin/nana /usr/local/bin/nana

ls -l --block-size=MB /usr/local/bin/nana

