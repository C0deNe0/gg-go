build:
	@go build -o bin/gg-go

run: build
	@./bin/gg-go

test:
	@go test -v ./...