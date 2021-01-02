.PHONY: build
build:
	go build ./cmd/server

.PHONY: run
run:
	go run ./cmd/server


.DEFAULT_GOAL: run