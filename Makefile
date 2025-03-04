all: run

run:
	go run cmd/app/main.go

test:
	go test -cover ./internal
