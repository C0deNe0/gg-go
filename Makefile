build:
	@go build -o bin/gg-go

go: build
	@./bin/gg-go

test:
	@go test -v ./...