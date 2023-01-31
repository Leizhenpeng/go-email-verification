database:
	docker-compose up -d
destroy:
	docker-compose down
dev:
	go run main.go

swag:
	swag init
	gow run ./

.PHONY: database dev destroy
