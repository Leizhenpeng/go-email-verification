database:
	docker-compose up -d
destroy:
	docker-compose down
dev:
	go run main.go

swag:
	swag init
	go run ./
.PHONY: database dev destroy
