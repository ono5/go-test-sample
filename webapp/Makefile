.PHONY: run test cover report

run:
	go run ./cmd/web/

test:
	go test -v ./...

cover:
	go test -cover ./...

report:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out
