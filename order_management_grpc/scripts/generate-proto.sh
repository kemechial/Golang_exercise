#!/bin/bash

# Script to generate Go code from proto files

echo "Generating Go code from proto files..."

# Create the output directory

# Generate the proto files into proto/orderpb
export PATH=$PATH:$HOME/go/bin
protoc --go_out=proto/orderpb --go_opt=paths=source_relative \
       --go-grpc_out=proto/orderpb --go-grpc_opt=paths=source_relative \
       proto/order.proto

echo "Proto generation completed!"
echo "Generated files:"
echo "  - proto/orderpb/order.pb.go"
echo "  - proto/orderpb/order_grpc.pb.go"
