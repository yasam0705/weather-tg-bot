-include .env
export

.PHONY: run
run:
	go run cmd/main.go

.PHONY: build-linux
build-linux:
	CGO_ENABLED=0 GOARCH="amd64" GOOS=linux go build -o ./bin/app cmd/main.go

.PHONY: migrate-up
migrate-up:
		migrate -source file://migrations -database postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DATABASE}?sslmode=disable up

.PHONY: migrate-down
migrate-down:
	migrate -source file://migrations -database postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DATABASE}?sslmode=disable down