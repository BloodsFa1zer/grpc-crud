# gRPC CRUD

## Description

This project is a gRPC-based CRUD service built with Go (Golang). It provides gRPC endpoints to perform CRUD operations on various resources. It leverages Protocol Buffers for message serialization and gRPC for communication.

## Features

- gRPC-based service
- CRUD operations for various resources
- Protocol Buffers for message serialization
- Integration with a PostgreSQL database
- Docker support for containerization

## Installation

### Prerequisites

- Go
- PostgreSQL
- Docker
- Protocol Buffers compiler (protoc)
- gRPC and Protocol Buffers Go plugins:
  ```sh
  go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
  go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

## Steps

1. Clone the repository:

    ```sh
    git clone https://github.com/BloodsFa1zer/grpc-crud.git
    cd grpc-crud
    ```

2. Install dependencies:

    ```sh
    go mod tidy
    ```

3. Set up the PostgreSQL database:
    - Create a database named `grpc_crud`.
    - Configure your environment variables as specified in the `env.dist` file.

4. Generate the gRPC code from the Protocol Buffers definitions:

    ```sh
    protoc --go_out=. --go-grpc_out=. proto/*.proto
    ```

5. Run the application:

    ```sh
    go run main.go
    ```

### Docker Setup

1. Build the Docker image:

    ```sh
    docker build -t grpc-crud .
    ```

2. Run the Docker container:

    ```sh
    docker run -p 50051:50051 grpc-crud
    ```

## Usage

### gRPC Endpoints

- **CreateResource** - Create a new resource
- **GetResource** - Get a resource by ID
- **UpdateResource** - Update a resource by ID
- **DeleteResource** - Delete a resource by ID
- **ListResources** - List all resources

### Example gRPC Client

You can use a gRPC client to interact with the service. Here is an example using the `grpcurl` command-line tool:

- List all resources:

    ```sh
    grpcurl -plaintext localhost:50051 list
    ```

- Create a new resource:

    ```sh
    grpcurl -plaintext -d '{"name": "New Resource"}' localhost:50051 ResourceService/CreateResource
    ```

- Get a resource by ID:

    ```sh
    grpcurl -plaintext -d '{"id": "1"}' localhost:50051 ResourceService/GetResource
    ```

## Configuration

The application configuration is managed through environment variables. Ensure that you set the appropriate values for your environment. You can refer to the `env.dist` file for the required variables:

```text
UserDatabasePassword=put_user_password_to_database_here
UserDatabaseName=put_user_name_to_database_here
DatabaseName=put_database_name_here
DriverDatabaseName=put_driver_database_name_here
GrpcAddr=put_your_server_addr_here
GrpcHost=put_your_server_host_here
