.PHONY: build
build:
		go build -v ./cmd/apiserver

.DEFAUL_GOAL := build