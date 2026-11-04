
DB_URL ?= "postgres://postgres:qwerty@localhost:5432/todo_db?sslmode=disable"

run:
	go run cmd/main.go

docker-up:
	docker compose up -d

docker-down:
	docker compose down

migrate-up:
	migrate -path ./schema -database $(DB_URL) up

migrate-down:
	migrate -path ./schema -database $(DB_URL) down