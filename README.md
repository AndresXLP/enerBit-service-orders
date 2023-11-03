``` sh
Clone with ssh recommended
$ git clone git@github.com:AndresXLP/enerBit-service-orders.git

Clone with https
$ git clone https://github.com/AndresXLP/enerBit-service-orders.git
```

# Requirements

* go v1.20
* go modules

# Technology Stack

- [GORM](https://gorm.io/)
- [gRPC](https://grpc.io/docs/languages/go/quickstart/)
- [Postgres](https://www.postgresql.org/docs/)
- [Redis Stream](https://redis.io/docs/data-types/streams/)

# Build

* Install dependencies:

```sh
$ go mod download
```

* [Migrations](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate) 
```sh
$ curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.2/migrate.linux-amd64.tar.gz | tar xvz
$ migrate -path ./migrations -database postgresql://${POSTGRESQL_DB_USER}:${POSTGRESQL_DB_PASSWORD}@${POSTGRESQL_DB_HOST}:${POSTGRESQL_DB_PORT}/${POSTGRESQL_DB_NAME}?sslmode=disable up
```

* Run local
```sh
$ go run cmd/main.go
```

# Environments

#### Required environment variables

* `GRPC_HOST`: host for gRPC server
* `GRPC_PORT`: port for gRPC server
* `GRPC_PROTOCOL`: protocol for gRPC server
* `POSTGRES_HOST`: host database
* `POSTGRES_USER`: user database
* `POSTGRES_PASSWORD`: password database
* `POSTGRES_DB_NAME`: name database
* `POSTGRES_PORT`: port database
* `REDIS_HOST`: host redis
* `REDIS_PORT`: port redis


# Contributors

* Andres Puello

