# someuser

This repository provides a basic example of a CRUD service(http, grpc) that interacts with various types of databases

## Table of Contents
- [Docker Installation](#Docker)
- [How It Works](#Jobs)
- [Stack](#Stack)
- [Struct project](#Struct)


## <a name="Docker"></a> Docker Installation
### Run
```
make up
```
### Volumes
[Configuration file](./config.yml) `/path/config.yml:/home/app/config.yml:ro`

[Environment file](./.env) `/path/.env:/home/app/.env:ro`

### Compose

The `docker-compose.yml` file contains all the necessary databases

## <a name="Jobs"></a> How It Works

**Note:** the `id` field has the format `any`, so in MongoDB it is `bson.ObjectId`, and in other databases it is `UUID`

The basic principle is the same for all databases

What you can do with a user:
- create user with validation
- get by id
- delete by id
- update by id with validation
- get all users

The `DB_TYPE` environment variable in the `./env` file allows you to switch between databases.

## <a name="Stack"></a> Stack

Backend
- **Golang**
- **Fiber**
- **Validator:**
- **goose:** migrations
- **gRPC**
- **swagger**

Data Base
- **PostgreSQL**
- **Redis**
- **MongoDB**
- **Cache**

**Note:** swagger documentation is available at `http://localhost:3000/swagger/index.html`

## <a name="Struct"></a> Struct project
```
someuser/
├── cmd                                # run app
│  └── up.go
├── docs                               # build swagger
│  ├── docs.go
│  ├── swagger.json
│  └── swagger.yaml
├── internal
│  ├── config
│  │  └── config.go
│  ├── model
│  │  ├── errormessage.go
│  │  └── someuser.go
│  ├── repository                      # db layer
│  │  ├── cache
│  │  │  └── cache.go
│  │  ├── factory
│  │  │  └── factory.go
│  │  ├── mongo
│  │  │  └── mongo.go
│  │  ├── postgres
│  │  │  └── postgres.go
│  │  ├── redis
│  │  │  └── redis.go
│  │  └── repository.go
│  ├── service                          # service layer
│  │  ├── factory
│  │  │  └── factory.go
│  │  ├── service.go
│  │  └── someuser
│  │      └── someuser.go
│  └── transport                        # transport layer
│      ├── grpc
│      │  ├── handler
│      │  │  └── someuser.go
│      │  └── proto
│      │      ├── someuser.pb.go
│      │      ├── someuser.proto
│      │      └── someuser_grpc.pb.go
│      └── http
│          └── handler
│              └── someuser.go
├── main.go
├── migrations
│  └── 20250411055852_init_someusers.sql   
└── pkg
    ├── logger
    │  └── logger.go
    └── utils
        └── utils.go                     # validation
```
