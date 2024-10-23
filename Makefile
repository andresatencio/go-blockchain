build:
	go build -o bin/go-blockchain

run: build
	bin/go-blockchain