.PHONY: build
build:
		sudo docker-compose build && sudo docker-compose up

.PHONY: migrate
migrate:
		migrate -path migrations -database postgres://api:password@127.0.0.1/api?sslmode=disable up

.PHONY:
drop:
		migrate -path migrations -database postgres://api:password@127.0.0.1/api?sslmode=disable drop

.DEFAULT_GOAL := build