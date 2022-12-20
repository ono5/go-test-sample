.PHONY: report cover tv ts

report:
	go test -coverprofile=coverage.out
	go tool cover -html=coverage.out

cover:
	go test -cover .

tv:
	go test -v .

ts:
	# make ts t=Test_isPrime
	# go test -v -run Test_alpha
	go test -v -run ${t}
