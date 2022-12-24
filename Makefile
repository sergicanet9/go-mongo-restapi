.PHONY: test

up:
	docker-compose up -d --build
down:
	docker-compose down
test-unit:
	go test -race $(shell go list ./... | grep -v /test) -coverprofile=coverage.out
	go tool cover -func=coverage.out -o=coverage.out
cover:
	go tool cover -html=coverage.out
test-integration:
	go test -race test/*.go
swagger:
	go install github.com/swaggo/swag/cmd/swag@v1.7.0
	swag init -g cmd/main.go -o app/docs
goose:
	go install github.com/pressly/goose/v3/cmd/goose@v3.5.0
	@read -p "Name for the change (e.g. add_column): " name; \
	goose -dir db/postgres/migrations/ create $${name:-<name>} sql
