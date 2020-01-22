.PHONY: build
build:
		sudo docker build -t apiserver . && sudo docker run --publish 8080:8080 --name apiserver --rm apiserver

.DEFAULT_GOAL := build