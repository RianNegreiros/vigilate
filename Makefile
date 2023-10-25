migrateup:
	migrate -path internal/database/migrations -database "postgresql://postgres:postgres@localhost:5432/postgres?sslmode=disable" -verbose up

migratedown:
	migrate -path internal/database/migrations -database "postgresql://postgres:postgres@localhost:5432/postgres?sslmode=disable" -verbose down

migration:
	@read -p "Enter migration name: " name; \
	migrate create -ext sql -dir internal/database/migrations $$name

dropdb:
	migrate -source file://internal/database/migrations -database "postgresql://postgres:postgres@localhost:5432/postgres?sslmode=disable" drop

run:
	go run cmd/main.go

.PHONY: run migrateup migratedown migration dropdb