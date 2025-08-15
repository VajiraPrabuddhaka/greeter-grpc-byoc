# gRPC Greeting Service

A simple gRPC service implementation in Go that provides greeting functionality with both unary and streaming RPC methods.

## Features

- **Unary RPC**: `SayHello` - Returns a simple greeting message
- **Server Streaming RPC**: `SayHelloStream` - Sends multiple greeting messages over time
- **Docker Support**: Multi-stage Docker build for containerized deployment
- **Build Automation**: Makefile with common development tasks

## Project Structure

```
.
├── proto/
│   └── greeting.proto          # gRPC service definition
├── client/
│   └── main.go                 # gRPC client implementation
├── scripts/
│   └── install-protoc.sh       # Script to install protoc plugins
├── main.go                     # gRPC server implementation
├── Dockerfile                  # Docker build configuration
├── Makefile                    # Build automation
├── go.mod                      # Go module definition
└── README.md                   # This file
```

## Prerequisites

- Go 1.24 or later
- Protocol Buffers compiler (protoc)
- Docker (optional, for containerized deployment)

## Setup

1. **Install protoc plugins** (if not already installed):
   ```bash
   ./scripts/install-protoc.sh
   ```

2. **Download dependencies**:
   ```bash
   make deps
   ```

## Building and Running

### Using Makefile

**Generate protobuf files and build binaries**:
```bash
make build
```

**Run the server**:
```bash
make run
```

**Run the client** (in another terminal):
```bash
make client
```

### Manual Commands

**Generate protobuf files**:
```bash
make proto
```

**Build server**:
```bash
go build -o bin/server .
```

**Build client**:
```bash
go build -o bin/client ./client
```

**Run server**:
```bash
./bin/server
```

**Run client**:
```bash
./bin/client
```

## Docker

**Build and run with Docker**:
```bash
make docker-run
```

**Or manually**:
```bash
docker build -t greeter-grpc .
docker run -p 50051:50051 greeter-grpc
```

## API Reference

### GreetingService

#### SayHello (Unary RPC)
- **Request**: `HelloRequest { name: string }`
- **Response**: `HelloResponse { message: string }`
- **Description**: Returns a single greeting message

#### SayHelloStream (Server Streaming RPC)
- **Request**: `HelloRequest { name: string }`
- **Response**: Stream of `HelloResponse { message: string }`
- **Description**: Returns 5 greeting messages with 1-second intervals

## Example Usage

When you run the client, you'll see output like:

```
Greeting: Hello, World!
Stream Greeting: Hello #1, Stream User!
Stream Greeting: Hello #2, Stream User!
Stream Greeting: Hello #3, Stream User!
Stream Greeting: Hello #4, Stream User!
Stream Greeting: Hello #5, Stream User!
```

## Development

**Clean build artifacts**:
```bash
make clean
```

**Install dependencies**:
```bash
make deps
```

## Configuration

- **Server Port**: 50051 (configurable in `main.go`)
- **Client Target**: localhost:50051 (configurable in `client/main.go`)

## License

This project is provided as-is for educational and demonstration purposes.