.PHONY: swag
swag:
	cd backend && go run github.com/swaggo/swag/cmd/swag@latest init -g main.go -o docs -d .

.PHONY: up
up:
	docker compose up --watch

.PHONY: up-build
up-build:
	docker compose up --build

.PHONY: down
down:
	docker compose down

.PHONY: test
test:
	docker compose exec api go test ./internal/tests/...

.PHONY: ci
ci:
	act -j ci
