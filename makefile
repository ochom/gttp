test:
	go test -v -coverprofile=coverage.out ./...

lint:
	golangci-lint run -v