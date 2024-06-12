# Makefile

.PHONY: run

run:
	@export $$(grep -v '^#' .env | xargs) && \
	go run ./cmd/main.go