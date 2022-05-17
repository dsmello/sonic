run:
	go run main.go

build:
	go build -a -o build/tb main.go

test:
	go test ./...