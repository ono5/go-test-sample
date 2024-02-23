.PHONY: run cover

run:
	go run ./cmd/web/

cover:
	go test ./cmd/web/ -coverprofile=coverage.out
	go tool cover -html=coverage.out
