## Payments Service

Go/Fiber microservice scaffold for handling payments-related flows (HTTP entrypoints currently stubbed). Includes gRPC modules, PostgreSQL via GORM, Redis caching, and multiple payment adapters.

### Requirements

- Go 1.22+ (module declares 1.25.4; use the latest available Go toolchain)
- Docker & Docker Compose
- Make (optional, for protobuf generation)

### Configuration

The service loads `config.yml` from the module root (working directory should be `cmd/server`, which resolves `../../config.yml`). Copy the example and adjust values:

```
cp config.example.yml config.yml
```

Key fields:

- `app.port`: HTTP port Fiber binds to (default `8080`)
- `db.*`: PostgreSQL connection info
- `redis.*`: Redis connection info

### Run locally (without Docker)

```
cp config.example.yml config.yml
cd cmd/server
go run .
```

### Docker

Build and run the container:

```
docker build -t payments-service .
docker run --rm -p 8080:8080 payments-service
```

### Docker Compose (recommended for local deps)

```
docker compose up --build
```

Services:

- `payments-service`: the app (exposes `8080`)
- `db`: PostgreSQL 16 (user/pass/db: `postgres/postgres/payments`)
- `redis`: Redis 7

### Protobuf generation

```
make generate-proto
```
