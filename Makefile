# runing docker-compose file for pgsql
postgres:
	docker compose up -d

# create posts database in postgres container
createdb:
	docker exec -it pgsql-db-1 createdb --username=postgres posts

# make init migration files
makeinitschema:
	migrate create -ext sql -dir db/migration -seq init_schema

# run migration file and create tables if write sql query in file
migrateup:
	migrate -path db/migration -database "postgresql://postgres:123@localhost:5433/posts?sslmode=disable" -verbose up

# drop all tables of posts database if exist
migratedown:
	migrate -path db/migration -database "postgresql://postgres:123@localhost:5433/posts?sslmode=disable" -verbose down

# delete posts database
dropdb:
	docker exec -it pgsql-db-1 dropdb posts

# docs in https://github.com/sqlc-dev/sqlc/tree/v1.4.0
sqlc:
	sqlc generate

test:
	go test -v -cover ./...

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test

# If you run above a command, you must run teminal `make command_name`. Example: make postgres