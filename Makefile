PWD := $(shell pwd)

.PHONY: run
run:
	go run ./cmd/worker


.PHONY: build
build:
	go build -o ./bin/dca-worker ./cmd/worker