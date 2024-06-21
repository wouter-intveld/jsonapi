build:
	@go build -o bin/jsonapi

run: build 
	@./bin/jsonapi

test:
	@go test -v ./...

startdb:
	@docker start jsonapi || docker run --name jsonapi -e POSTGRES_PASSWORD=$POSTGRES_PASSWORD -p 5432:5432 -d postgres

stopdb:
	@docker stop jsonapi

