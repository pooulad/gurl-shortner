build:
	@go build -o ./bin/gurl-shortner ./cmd/app/main.go

run: build
	@./bin/gurl-shortner

tidy:
	@go mod tidy

testgo:
	@go test ./... -v