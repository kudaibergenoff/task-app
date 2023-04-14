.PHONY:

run:
	@echo Run go app
	go run cmd/app/main.go
dev:
	@echo docker compose up build
	docker compose -f docker-compose.yml up --build