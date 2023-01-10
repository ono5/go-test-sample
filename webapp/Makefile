.PHONY: run test cover report ts api curl

run:
	go run ./cmd/web/

api:
	go run ./cmd/api/

test:
	# go test -v ./...
	go test -v -tags=integration ./...

cover:
	go test -cover ./...

report:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out

ts:
	# make ts t=Test_isPrime
	# go test -v -run Test_alpha
	go test -v -run ${t} ./...

curl:
	curl http://localhost:8090/test
