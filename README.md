# someuser

This repository provides a basic example of a CRUD service that interacts with various types of databases:
- PostgreSQL
- MongoDB
- Redis
- Cache


## Table of Contents
- [Docker Installation](#Docker)
- [How It Works](#Jobs)


## <a name="Docker"></a> Docker Installation
### Image
```
docker compose up --build -d
```
### Volumes
[Configuration file](./config.yml) `/path/config.yml:/home/app/config.yml:ro`

[Environment file](./.env) `/path/.env:/home/app/.env:ro`

### Compose

The `docker-compose.yml` file contains all the necessary databases

## <a name="Jobs"></a> How It Works

The basic principle is the same for all databases

What you can do with a user:
- create
- get by id
- delete
- update
- get all users

The `DB_TYPE` environment variable in the `./env` file allows you to switch between databases.

**Note:** swagger documentation is available at `http://localhost:3000/swagger/index.html`