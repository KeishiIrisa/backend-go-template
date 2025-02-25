.PHONY: swag
swag:
	go run github.com/swaggo/swag/cmd/swag@latest init -g backend/main.go -o backend/docs

.PHONY: up
up:
	docker-compose up --watch

.PHONY: up-build
up-build:
	docker-compose up --build

.PHONY: down
down:
	docker-compose down

.PHONY: test
test:
	docker-compose exec api go test ./internal/tests/...
