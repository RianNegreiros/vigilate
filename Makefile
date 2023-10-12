migrateup:
	migrate -path db/migrations -database "postgresql://postgres:postgres@localhost:5432/postgres?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migrations -database "postgresql://postgres:postgres@localhost:5432/postgres?sslmode=disable" -verbose down

migration:
	@read -p "Enter migration name: " name; \
	migrate create -ext sql -dir db/migrations $$name

dropdb:
	migrate -source file://db/migrations -database "postgresql://postgres:postgres@localhost:5432/postgres?sslmode=disable" drop

run:
	go run cmd/main.go

.PHONY: run migrateup migratedown