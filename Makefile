cover:
	go test ./... -coverprofile=coverage.out && go tool cover -html=coverage.out && rm coverage.out

migrate:
	go run cmd/migrate.go