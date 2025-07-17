.PHONY: build run stop clean test

build:
	docker-compose build

run:
	docker-compose up -d

stop:
	docker-compose down

clean:
	docker-compose down -v
	docker system prune -f

logs:
	docker-compose logs -f

test:
	cd backend && go test ./...

dev:
	cd backend && go run cmd/main.go 