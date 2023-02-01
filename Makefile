db_up:
	docker-compose up -d
db_down:
	docker-compose down
dev:
	go run main.go

swag:
	swag init
	gow run ./

.PHONY: db_up db_down destroy
