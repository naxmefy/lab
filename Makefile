.PHONY: run
run:
	@go run cmd/lab/main.go

.PHONY: test
test:
	@go test -cover ./...

.PHONY: install
install:
	@go install ./...

.PHONY: build
build:
	@go build -o bin/lab
