include .env

.PHONY: fmt
fmt:
	goimports -l -w internal/

web:
	go run ./cmd/web
