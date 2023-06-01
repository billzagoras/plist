appserver:
	go build -o appserver ./cmd

test:
	go test -v ./...