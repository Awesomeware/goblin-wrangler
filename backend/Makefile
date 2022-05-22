include .env

.PHONY: fmt
fmt:
	goimports -l -w internal/

backend:
	go run ./cmd/backend
