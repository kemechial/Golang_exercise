#!/bin/bash

# Script to generate Go code from proto files

echo "Generating Go code from proto files..."

export PATH=$PATH:$HOME/go/bin

# Generate order service proto files
protoc --go_out=proto/orderpb --go_opt=paths=source_relative \
       --go-grpc_out=proto/orderpb --go-grpc_opt=paths=source_relative \
       proto/order.proto

# Generate inventory service proto files
protoc --go_out=proto/inventorypb --go_opt=paths=source_relative \
       --go-grpc_out=proto/inventorypb --go-grpc_opt=paths=source_relative \
       proto/inventory.proto

# Generate delivery service proto files
protoc --go_out=proto/deliverypb --go_opt=paths=source_relative \
       --go-grpc_out=proto/deliverypb --go-grpc_opt=paths=source_relative \
       proto/delivery.proto

echo "Proto generation completed!"
echo "Generated files:"
echo "  - proto/orderpb/order.pb.go"
echo "  - proto/orderpb/order_grpc.pb.go"
echo "  - proto/inventorypb/inventory.pb.go"
echo "  - proto/inventorypb/inventory_grpc.pb.go"
echo "  - proto/deliverypb/delivery.pb.go"
echo "  - proto/deliverypb/delivery_grpc.pb.go"
