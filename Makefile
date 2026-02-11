lb:
	@go run cmd/lb/main.go 

stress:
	@go run cmd/stress-test/main.go -port ":80"

test:
	@go tool staticcheck .
	@go test -v ./...

fuzz:
	@go test -run NONE -fuzz . -fuzztime 10s
