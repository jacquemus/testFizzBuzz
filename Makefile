## test: run unit and integration tests
test: test-integration test-unit

## test-unit: run unit tests
test-unit:
	go clean -testcache
	go test ./fizzer -v -cover

## test-integration: run integration tests
test-integration:
	go clean -testcache
	go test . -v -cover

## update: runs go mod vendor and tidy
update: mod tidy

## build: build the microservices
build: go build .

## mod: runs go mod vendor
mod:
	go mod vendor -v

## tidy: runs go mod tidy
tidy:
	go mod tidy -v

lint:
	golangci-lint --version
	golangci-lint run -v ./...
