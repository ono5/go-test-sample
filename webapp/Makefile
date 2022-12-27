.PHONY: run test cover report ts

run:
	go run ./cmd/web/

test:
	go test -v ./...

cover:
	go test -cover ./...

report:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out

ts:
	# make ts t=Test_isPrime
	# go test -v -run Test_alpha
	go test -v -run ${t} ./...
