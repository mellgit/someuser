# someuser

This repository provides a basic example of a CRUD service(http, grpc) that interacts with various types of databases

## Table of Contents
- [Docker Installation](#Docker)
- [How It Works](#Jobs)
- [Stack](#Stack)


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

Data Base
- **PostgreSQL**
- **Redis**
- **MongoDB**
- **Cache**

**Note:** swagger documentation is available at `http://localhost:3000/swagger/index.html`
