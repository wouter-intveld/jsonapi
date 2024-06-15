build:
	@go build -o bin/jsonapi

run: build
	@./bin/jsonapi

test:
	@go test -v ./...


