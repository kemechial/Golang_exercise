#!/bin/bash

# Script to generate Go code from proto files

echo "Generating Go code from proto files..."

# Create the output directory
mkdir -p proto/orderpb

# Generate the proto files
export PATH=$PATH:$HOME/go/bin
protoc --go_out=. --go_opt=paths=source_relative \
       --go-grpc_out=. --go-grpc_opt=paths=source_relative \
       proto/order.proto

echo "Proto generation completed!"
echo "Generated files:"
echo "  - proto/orderpb/order.pb.go"
echo "  - proto/orderpb/order_grpc.pb.go"
