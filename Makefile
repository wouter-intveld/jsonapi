build:
	@go build -o bin/jsonapi

run: build
	@./bin/jsonapi

test:
	@go test -v ./...

startdb:
	@docker start jsonapi

stopdb:
	@docker stop jsonapi

