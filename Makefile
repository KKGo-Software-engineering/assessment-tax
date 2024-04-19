run:
	PORT=8080 DATABASE_URL="host=localhost port=5432 user=postgres password=postgres dbname=ktaxes sslmode=disable" go run main.go

test: 
	go test -v ./...

test-cover:
	go test -coverprofile coverage.html ./...
	go tool cover -html=coverage.html