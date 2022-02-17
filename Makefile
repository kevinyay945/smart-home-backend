.PHONY: generate run migrate

run:
	go run .

generate:
	go generate ./...

migrate:
	go run lib/pg/migration/main.go