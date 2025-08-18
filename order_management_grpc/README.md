# Order Management gRPC Service

A comprehensive gRPC-based order management system built in Go. This project demonstrates advanced gRPC concepts including unary calls, server streaming, client streaming, and bidirectional streaming.

## 🚀 Features

- **Complete Order Lifecycle**: Create, read, update, and cancel orders
- **Real-time Updates**: Watch order status changes using server streaming
- **Batch Operations**: Create multiple orders using client streaming
- **Bidirectional Streaming**: Real-time order processing
- **Status Validation**: Enforces valid order status transitions
- **In-memory Storage**: Fast, simple storage for development and testing
- **Comprehensive Testing**: Includes client for testing all features

## 📋 Order States

The system enforces the following order status transitions:

```
PENDING → CONFIRMED → PREPARING → SHIPPED → DELIVERED
    ↓         ↓          ↓
CANCELLED  CANCELLED  CANCELLED
```

## 🏗️ Project Structure

```
order_management_grpc/
├── cmd/
│   ├── server/          # gRPC server implementation
│   └── client/          # Test client implementation
├── internal/
│   ├── models/          # Business models and conversions
│   ├── service/         # gRPC service implementation
│   └── storage/         # Storage layer (in-memory)
├── proto/               # Protocol buffer definitions
├── scripts/             # Build and utility scripts
├── Makefile            # Build automation
└── README.md           # This file
```

## 🛠️ Prerequisites

- Go 1.21 or later
- Protocol Buffers compiler (`protoc`)
- Go protobuf plugins

## 📦 Installation

1. **Clone the repository** (or create in your workspace):
   ```bash
   cd /home/kaan/Golang_exercise/order_management_grpc
   ```

2. **Install protoc plugins**:
   ```bash
   make install-protoc
   ```

3. **Install dependencies and tidy modules**:
   ```bash
   make install-deps
   ```

4. **Generate proto files**:
   ```bash
   make proto
   ```

## 🎯 Quick Start

### Option 1: Using Make (Recommended)

1. **Start the server**:
   ```bash
   make server
   ```

2. **In another terminal, run the client**:
   ```bash
   make client
   ```

### Option 2: Manual Build and Run

1. **Build the project**:
   ```bash
   make build
   ```

2. **Run the server**:
   ```bash
   ./bin/server
   ```

3. **Run the client** (in another terminal):
   ```bash
   ./bin/client
   ```

## 🔧 API Reference

### Unary RPC Methods

#### CreateOrder
Creates a new order with specified items.

```protobuf
rpc CreateOrder(CreateOrderRequest) returns (CreateOrderResponse);
```

#### GetOrder
Retrieves an order by its ID.

```protobuf
rpc GetOrder(GetOrderRequest) returns (GetOrderResponse);
```

#### UpdateOrderStatus
Updates the status of an existing order.

```protobuf
rpc UpdateOrderStatus(UpdateOrderStatusRequest) returns (UpdateOrderStatusResponse);
```

#### ListOrders
Lists orders with optional customer filtering and pagination.

```protobuf
rpc ListOrders(ListOrdersRequest) returns (ListOrdersResponse);
```

#### CancelOrder
Cancels an order (if in a cancellable state).

```protobuf
rpc CancelOrder(CancelOrderRequest) returns (CancelOrderResponse);
```

### Streaming RPC Methods

#### WatchOrderStatus (Server Streaming)
Streams real-time updates for a specific order.

```protobuf
rpc WatchOrderStatus(GetOrderRequest) returns (stream Order);
```

#### BatchCreateOrders (Client Streaming)
Accepts a stream of order creation requests and returns all created orders.

```protobuf
rpc BatchCreateOrders(stream CreateOrderRequest) returns (ListOrdersResponse);
```

#### ProcessOrdersStream (Bidirectional Streaming)
Processes order status updates in real-time, both ways.

```protobuf
rpc ProcessOrdersStream(stream UpdateOrderStatusRequest) returns (stream Order);
```

## 🧪 Testing with grpcurl

If you have `grpcurl` installed, you can test the API directly:

```bash
# List available services
grpcurl -plaintext localhost:50051 list

# List methods for OrderService
grpcurl -plaintext localhost:50051 list order.OrderService

# Create an order
grpcurl -plaintext -d '{
  "customer_id": "customer-123",
  "customer_name": "John Doe",
  "items": [{
    "product_id": "product-1",
    "product_name": "Laptop",
    "quantity": 1,
    "unit_price": 999.99
  }]
}' localhost:50051 order.OrderService/CreateOrder

# Get an order (replace ORDER_ID with actual ID)
grpcurl -plaintext -d '{"order_id": "ORDER_ID"}' localhost:50051 order.OrderService/GetOrder
```

## 📝 Learning Objectives

This project covers the following gRPC and Go concepts:

### gRPC Concepts
- [x] **Unary RPCs**: Simple request-response pattern
- [x] **Server Streaming**: Server sends multiple responses
- [x] **Client Streaming**: Client sends multiple requests
- [x] **Bidirectional Streaming**: Both client and server stream
- [x] **Error Handling**: Proper gRPC status codes
- [x] **Proto3 Syntax**: Modern protobuf definition
- [x] **Service Reflection**: For debugging and tooling

### Go Concepts
- [x] **Interface Design**: Clean separation of concerns
- [x] **Concurrent Programming**: Safe goroutine communication
- [x] **Context Usage**: Proper request lifecycle management
- [x] **Error Handling**: Comprehensive error management
- [x] **Package Organization**: Well-structured Go project
- [x] **Testing**: Client-based integration testing
- [x] **Build Automation**: Makefile for common tasks

### Software Engineering Practices
- [x] **Domain Modeling**: Proper business logic modeling
- [x] **Status Transitions**: State machine implementation
- [x] **Data Validation**: Input validation and sanitization
- [x] **Real-time Communication**: Event-driven architecture
- [x] **API Design**: RESTful-like gRPC API design

## 🎓 Next Steps

To further enhance your learning:

1. **Add Persistence**: Replace in-memory storage with a database
2. **Add Authentication**: Implement JWT or other auth mechanisms
3. **Add Metrics**: Integrate Prometheus metrics
4. **Add Logging**: Structured logging with context
5. **Add Docker**: Containerize the application
6. **Add Tests**: Unit and integration tests
7. **Add Validation**: More comprehensive input validation
8. **Add Documentation**: Generate API documentation from proto files

## 🛠️ Development Commands

```bash
# Format code
make fmt

# Run tests
make test

# Clean build artifacts
make clean

# View all available commands
make help
```

## 📄 License

This project is for educational purposes. Feel free to use and modify as needed for learning gRPC and Go.

## 🤝 Contributing

This is a learning project, but suggestions and improvements are welcome! Feel free to:
- Report issues
- Suggest improvements
- Add new features
- Improve documentation

---

**Happy Learning!** 🚀

This project provides a solid foundation for understanding gRPC in Go. Start with the basic unary calls, then explore the streaming capabilities to see the full power of gRPC.
