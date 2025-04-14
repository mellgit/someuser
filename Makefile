include .env
LOCAL_BIN:=$(PWD)/bin

install-deps:
	GOBIN=$(LOCAL_BIN) go install github.com/pressly/goose/v3/cmd/goose@v3.23.0
	#GOBIN=$(LOCAL_BIN) go install github.com/sqlc-dev/sqlc/cmd/sqlc@v1.28.0
	GOBIN=$(LOCAL_BIN) go install github.com/swaggo/swag/cmd/swag@v1.16.3

migration-status:
	$(LOCAL_BIN)/goose -dir $(POSTGRES_MIGRATIONS_PATH) postgres $(POSTGRES_MIGRATIONS_DSN) status -v
migration-add:
	$(LOCAL_BIN)/goose -dir $(POSTGRES_MIGRATIONS_PATH) create $(name) sql
migration-up:
	$(LOCAL_BIN)/goose -dir $(POSTGRES_MIGRATIONS_PATH) postgres $(POSTGRES_MIGRATIONS_DSN) up -v
migration-down:
	$(LOCAL_BIN)/goose -dir $(POSTGRES_MIGRATIONS_PATH) postgres $(POSTGRES_MIGRATIONS_DSN) down -v

gen-sql:
	$(LOCAL_BIN)/sqlc generate

swag:
	$(LOCAL_BIN)/swag init -g cmd/up.go

# docker basic
pa:
	docker ps -a
up:
	docker compose up --build -d
down:
	docker compose down
i:
	docker images
b:
	docker build -t someuser .
cleardb:
	rm -r ./mongo_data ./postgres_data
r: down up
cn:
	docker rmi $(docker images -q --filter="dangling=true") # clear images with none status


lu:
	docker logs -f --tail 100 someuser
eu:
	docker exec -it someuser sh
