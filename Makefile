.PHONY: build
build:
		sudo docker-compose build && sudo docker-compose up

.PHONY: migrate
migrate:
		migrate -path migrations -database postgres://api:password@127.0.0.1/api?sslmode=disable up

.PHONY: drop-db
drop-db:
		migrate -path migrations -database postgres://api:password@127.0.0.1/api?sslmode=disable drop

.PHONY: drop-docker
drop-docker:
		sudo docker system prune -a --volumes

.DEFAULT_GOAL := build