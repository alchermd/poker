up:
	go build -o bin/server ./cmd/server
	./bin/server

clean:
	rm -rf ./bin

format:
	go fmt ./...
