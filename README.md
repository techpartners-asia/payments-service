# Payments Service

A production-ready microservice for handling payment processing with support for multiple payment gateways, merchant management, and invoice generation. Built with Go, Fiber, gRPC, PostgreSQL, and Redis.

## Table of Contents

- [Overview](#overview)
- [Architecture](#architecture)
- [Features](#features)
- [Supported Payment Gateways](#supported-payment-gateways)
- [Requirements](#requirements)
- [Configuration](#configuration)
- [Installation & Setup](#installation--setup)
- [Development](#development)
- [Docker Deployment](#docker-deployment)
- [API Documentation](#api-documentation)
- [Database Schema](#database-schema)
- [Project Structure](#project-structure)
- [Contributing](#contributing)

## Overview

The Payments Service is a microservice designed to handle payment processing workflows including:

- **Payment Creation**: Generate invoices through multiple payment gateway adapters
- **Payment Verification**: Check payment status and handle callbacks
- **Merchant Management**: Manage merchant accounts with payment gateway credentials
- **Ebarimt Integration**: Handle tax invoice (ebarimt) credentials for merchants
- **Caching**: Redis-based caching for improved performance
- **Multi-Gateway Support**: Unified interface for 9+ payment providers

## Architecture

The service follows a clean architecture pattern with clear separation of concerns:

```
┌─────────────────────────────────────────────────────────┐
│                    API Layer                             │
│  ┌──────────────┐              ┌──────────────┐         │
│  │ HTTP (Fiber) │              │   gRPC       │         │
│  │   Port 8080  │              │  Port 50051 │         │
│  └──────────────┘              └──────────────┘         │
└─────────────────────────────────────────────────────────┘
                        │
┌─────────────────────────────────────────────────────────┐
│                 Business Logic Layer                     │
│  ┌──────────────┐              ┌──────────────┐         │
│  │   Payment    │              │   Merchant   │         │
│  │   Use Cases  │              │   Use Cases  │         │
│  └──────────────┘              └──────────────┘         │
└─────────────────────────────────────────────────────────┘
                        │
┌─────────────────────────────────────────────────────────┐
│              Infrastructure Layer                       │
│  ┌──────────────┐  ┌──────────────┐  ┌──────────────┐ │
│  │  PostgreSQL  │  │    Redis     │  │   Payment    │ │
│  │   (GORM)     │  │   (Cache)    │  │   Adapters   │ │
│  └──────────────┘  └──────────────┘  └──────────────┘ │
└─────────────────────────────────────────────────────────┘
```

### Technology Stack

- **Language**: Go 1.25.4
- **Web Framework**: Fiber v2 (HTTP server)
- **RPC Framework**: gRPC (for inter-service communication)
- **ORM**: GORM (PostgreSQL driver)
- **Cache**: Redis (go-redis v9)
- **Configuration**: Viper (YAML-based)
- **Error Tracking**: Sentry
- **JSON Processing**: Sonic (high-performance)

## Features

### Core Functionality

- ✅ **Multi-Gateway Payment Processing**: Unified interface for 9+ payment providers
- ✅ **Payment Lifecycle Management**: Create, check, and track payment status
- ✅ **Merchant Management**: CRUD operations for merchants with gateway credentials
- ✅ **Ebarimt Integration**: Tax invoice credential management
- ✅ **Redis Caching**: Performance optimization with caching layer
- ✅ **gRPC & HTTP APIs**: Dual interface for different use cases
- ✅ **Payment Logging**: Comprehensive audit trail for all transactions
- ✅ **Streaming Support**: Real-time payment status updates via gRPC streams

### Payment Status Flow

```
Pending → Paid | Cancelled | Refunded
```

- **Pending**: Payment invoice created, awaiting customer payment
- **Paid**: Payment successfully completed
- **Cancelled**: Payment cancelled by customer or merchant
- **Refunded**: Payment refunded to customer

## Supported Payment Gateways

The service supports the following payment gateway adapters:

| Gateway              | Type        | Status    | Description                       |
| -------------------- | ----------- | --------- | --------------------------------- |
| **QPay**             | `qpay`      | ✅ Active | QPay payment gateway integration  |
| **TokiPay**          | `tokipay`   | ✅ Active | TokiPay mobile payment solution   |
| **Monpay**           | `monpay`    | ✅ Active | Monpay payment gateway            |
| **Golomt Ecommerce** | `ecommerce` | ✅ Active | Golomt Bank ecommerce integration |
| **SocialPay**        | `socialpay` | ✅ Active | SocialPay payment gateway         |
| **StorePay**         | `storepay`  | ✅ Active | StorePay payment solution         |
| **Pocket**           | `pocket`    | ✅ Active | Pocket payment gateway            |
| **Simple**           | `simple`    | ✅ Active | Simple payment gateway            |
| **Balc**             | `balc`      | ✅ Active | Balc credit card integration      |

Each adapter implements a standardized interface for:

- **CreateInvoice**: Generate payment invoices
- **CheckInvoice**: Verify payment status

## Requirements

### Runtime Requirements

- **Go**: 1.22+ (module declares 1.25.4)
- **PostgreSQL**: 12+ (16 recommended)
- **Redis**: 6+ (7 recommended)
- **Docker**: 20.10+ (for containerized deployment)
- **Docker Compose**: 2.0+ (for local development)

### Development Requirements

- **Make**: For protobuf code generation
- **Protocol Buffers Compiler**: `protoc` with Go plugins
- **Git**: For version control

## Configuration

The service uses YAML-based configuration loaded from `config.yml` at the project root. The configuration is resolved relative to the working directory (`cmd/server` resolves `../../config.yml`).

### Configuration File Structure

```yaml
app:
  name: payments-service # Service name
  version: "0.1.0" # Service version
  port: "8080" # HTTP server port
  env: development # Environment (development/staging/production)

db:
  host: db # PostgreSQL host
  port: "5432" # PostgreSQL port
  user: postgres # Database user
  name: payments # Database name
  password: postgres # Database password
  timezone: Asia/Ulaanbaatar # Database timezone

redis:
  host: redis # Redis host
  port: "6379" # Redis port
  password: "" # Redis password (empty for no auth)
  db: 0 # Redis database number
```

### Environment-Specific Configuration

1. **Development**: Use `config.example.yml` as a template
2. **Production**: Override via environment variables or mounted config files
3. **Docker**: Mount custom config file or use environment variables

### Payment Gateway Configuration

Payment gateway credentials are managed per-merchant through the gRPC API. Each merchant can configure multiple payment gateways with their respective credentials (API keys, endpoints, callbacks, etc.).

## Installation & Setup

### 1. Clone the Repository

```bash
git clone <repository-url>
cd payments-service
```

### 2. Install Dependencies

```bash
go mod download
```

### 3. Generate Protobuf Code

```bash
make generate-proto
```

This generates gRPC service code from `.proto` files in `pkg/proto/`.

### 4. Configure the Service

```bash
cp config.example.yml config.yml
# Edit config.yml with your database and Redis credentials
```

### 5. Database Setup

The service uses GORM for database migrations. Ensure PostgreSQL is running and the database specified in `config.yml` exists:

```sql
CREATE DATABASE payments;
```

The service will automatically create tables on first run based on GORM entity definitions.

## Development

### Running Locally (Without Docker)

#### Prerequisites

1. PostgreSQL running locally or remotely
2. Redis running locally or remotely
3. Configuration file set up

#### Start the HTTP Server

```bash
cd cmd/server
go run .
```

The HTTP server will start on port `8080` (or as configured).

#### Start the gRPC Server

```bash
cd cmd/grpc
go run .
```

The gRPC server will start on port `50051`.

### Running with Docker Compose (Recommended)

Docker Compose runs the payment service containers. You need to provide your own PostgreSQL and Redis instances (configured in `config.yml`):

```bash
# First, create config.yml from the example and configure your database/Redis
cp config.example.yml config.yml
# Edit config.yml with your database and Redis connection details

# Start the services
docker compose up --build
```

This starts:

- **payments-service-http**: HTTP server (Fiber) on port 8080
- **payments-service-grpc**: gRPC server on port 50051

Both HTTP and gRPC servers run as separate containers for better scalability and resource management.

**Important**:

- The service connects to external PostgreSQL and Redis instances specified in your `config.yml`
- Make sure your database and Redis are accessible from the containers
- Update `config.yml` with your actual database host, port, credentials, and Redis connection details

### Hot Reloading (Development)

For development with hot reloading, use tools like:

- **Air**: `go install github.com/cosmtrek/air@latest`
- **CompileDaemon**: `go install github.com/githubnemo/CompileDaemon@latest`

Example with Air:

```bash
# Install Air
go install github.com/cosmtrek/air@latest

# Run with hot reload
cd cmd/server
air
```

## Docker Deployment

### Building the Docker Image

```bash
docker build -t payments-service:latest .
```

### Running the Container

The Dockerfile builds both HTTP and gRPC server binaries. By default, the container runs the HTTP server.

**HTTP Server:**

```bash
docker run -d \
  --name payments-service-http \
  -p 8080:8080 \
  -v $(pwd)/config.yml:/app/config.yml:ro \
  payments-service:latest
```

**gRPC Server:**

```bash
docker run -d \
  --name payments-service-grpc \
  -p 50051:50051 \
  -v $(pwd)/config.yml:/app/config.yml:ro \
  --entrypoint /usr/local/bin/payments-service-grpc \
  payments-service:latest
```

**Note**: For production, it's recommended to run both services as separate containers (as shown in docker-compose.yml) for better resource management and scalability.

### Docker Compose for Production

The docker-compose.yml is configured to connect to external databases. For production:

1. Ensure your PostgreSQL and Redis instances are running and accessible
2. Configure `config.yml` with production database credentials
3. Use environment variables or secrets management for sensitive data
4. Start the services:

```bash
docker compose up -d
```

For environment-specific configurations, you can use multiple compose files:

```bash
docker compose -f docker-compose.yml -f docker-compose.prod.yml up -d
```

## API Documentation

### gRPC API

The service exposes two gRPC services:

#### PaymentService

**Service**: `PaymentService`

**Methods**:

1. **Create** - Create a new payment invoice

   ```protobuf
   rpc Create(PaymentCreateRequest) returns (PaymentCreateResponse);
   ```

   - Request: `amount` (float)
   - Response: `uid`, `amount`, `status`, `invoiceID`

2. **Check** - Check payment status

   ```protobuf
   rpc Check(PaymentCheckRequest) returns (PaymentCheckResponse);
   ```

   - Request: `uid` (string)
   - Response: `uid`, `amount`, `status`

3. **CheckStream** - Stream payment status updates
   ```protobuf
   rpc CheckStream(PaymentCheckRequest) returns (stream PaymentCheckResponse);
   ```
   - Real-time status updates via gRPC streaming

**Payment Status Enum**:

- `PAYMENT_STATUS_PENDING` (0)
- `PAYMENT_STATUS_PAID` (1)
- `PAYMENT_STATUS_CANCELLED` (2)
- `PAYMENT_STATUS_REFUNDED` (3)

#### MerchantService

**Service**: `MerchantService`

**Methods**:

1. **Create** - Create a new merchant

   ```protobuf
   rpc Create(CreateMerchantRequest) returns (MerchantResponse);
   ```

2. **GetByID** - Get merchant by ID

   ```protobuf
   rpc GetByID(MerchantIDRequest) returns (MerchantResponse);
   ```

3. **Update** - Update merchant information

   ```protobuf
   rpc Update(UpdateRequest) returns (MerchantResponse);
   ```

4. **Delete** - Delete a merchant
   ```protobuf
   rpc Delete(MerchantIDRequest) returns (SuccessResponse);
   ```

### HTTP API

The HTTP API is built with Fiber and currently serves as a REST interface (routes can be extended in `internal/delivery/http/routes/base.go`).

**Base URL**: `http://localhost:8080`

### Example gRPC Client Usage

```go
import (
    "context"
    "google.golang.org/grpc"
    paymentProto "git.techpartners.asia/gateway-services/payment-service/pkg/proto/payment"
)

conn, _ := grpc.Dial("localhost:50051", grpc.WithInsecure())
client := paymentProto.NewPaymentServiceClient(conn)

// Create payment
resp, _ := client.Create(context.Background(), &paymentProto.PaymentCreateRequest{
    Amount: 10000.0,
})

// Check payment
status, _ := client.Check(context.Background(), &paymentProto.PaymentCheckRequest{
    Uid: resp.Uid,
})
```

## Database Schema

### Core Entities

#### PaymentEntity

Stores payment transaction information:

- `id`: Primary key (auto-increment)
- `uid`: Unique payment identifier (indexed)
- `status`: Payment status (pending/paid/cancelled/refunded)
- `amount`: Payment amount
- `phone`: Customer phone number
- `customer_id`: Customer identifier (indexed)
- `note`: Payment notes
- `ref_invoice_id`: Reference invoice ID from payment gateway
- `merchant_id`: Associated merchant (indexed, foreign key)
- `type`: Payment gateway type
- `created_at`: Creation timestamp
- `updated_at`: Last update timestamp

#### PaymentLogEntity

Audit trail for payment operations:

- `id`: Primary key
- `payment_id`: Associated payment (indexed, foreign key)
- `message`: Log message
- `created_at`: Log timestamp

#### MerchantEntity

Merchant account information:

- `id`: Primary key
- `name`: Merchant name
- `created_at`: Creation timestamp
- `updated_at`: Last update timestamp

#### MerchantEbarimtEntity

Tax invoice (ebarimt) credentials for merchants:

- `id`: Primary key
- `merchant_id`: Associated merchant (foreign key)
- `url`: Ebarimt service URL
- `tin`: Tax identification number
- `pos_no`: Point of sale number
- `branch_no`: Branch number
- `district_code`: District code

### Relationships

```
MerchantEntity (1) ──< (N) PaymentEntity
PaymentEntity (1) ──< (N) PaymentLogEntity
MerchantEntity (1) ──< (1) MerchantEbarimtEntity
```

## Project Structure

```
payments-service/
├── cmd/
│   ├── grpc/              # gRPC server entry point
│   │   └── main.go
│   └── server/            # HTTP server entry point
│       └── main.go
├── infrastructure/
│   ├── database/          # Database layer (GORM)
│   │   ├── entity/        # Database entities
│   │   └── repository/    # Data access layer
│   ├── payment/            # Payment gateway adapters
│   │   └── adapters/      # Gateway-specific implementations
│   ├── redis/             # Redis caching layer
│   └── shared/            # Shared infrastructure code
├── internal/
│   ├── delivery/          # API delivery layer
│   │   ├── grpc/          # gRPC handlers
│   │   └── http/          # HTTP handlers (Fiber)
│   └── modules/           # Business logic modules
│       ├── merchant/      # Merchant management
│       └── payment/       # Payment processing
├── pkg/
│   ├── config/            # Configuration management
│   ├── fiber/             # Fiber framework setup
│   ├── proto/             # Protobuf definitions
│   ├── sentry/            # Error tracking
│   └── utils/             # Utility functions
├── config.example.yml     # Configuration template
├── docker-compose.yml     # Docker Compose setup
├── Dockerfile             # Docker image definition
├── go.mod                 # Go module dependencies
├── Makefile               # Build automation
└── README.md              # This file
```

## Contributing

### Development Workflow

1. Create a feature branch from `main`
2. Make your changes
3. Run tests (if available)
4. Update documentation
5. Submit a pull request

### Code Style

- Follow Go standard formatting (`gofmt`)
- Use meaningful variable and function names
- Add comments for exported functions and types
- Keep functions focused and small

### Adding a New Payment Gateway

1. Create adapter in `infrastructure/payment/adapters/`
2. Implement `CreateInvoice` and `CheckInvoice` methods
3. Add payment type constant in `infrastructure/database/entity/merchant.go`
4. Register adapter in `infrastructure/payment/base.go`
5. Update protobuf definitions if needed
6. Update this README

## License

See [LICENSE](LICENSE) file for details.

## Support

For issues, questions, or contributions, please open an issue on the repository.

---

**Built with ❤️ using Go, Fiber, and gRPC**
