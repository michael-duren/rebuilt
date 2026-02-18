servers:
	@for port in 8081 8082 8083; do \
		if ! curl -s http://localhost:$$port > /dev/null 2>&1; then \
			echo "Starting server on port $$port..."; \
			docker start server-$$port 2>/dev/null || docker run -d --name server-$$port -p $$port:80 lb-test-server; \
		fi \
	done

lb:
	@go run cmd/lb/main.go 

stress:
	@go run cmd/stress-test/main.go -port ":80"

test:
	@go tool staticcheck .
	@go test -v ./...

fuzz:
	@go test -run NONE -fuzz . -fuzztime 10s
