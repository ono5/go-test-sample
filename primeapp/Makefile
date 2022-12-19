.PHONY: report cover

report:
	go test -coverprofile=coverage.out
	go tool cover -html=coverage.out

cover:
	go test -cover .
