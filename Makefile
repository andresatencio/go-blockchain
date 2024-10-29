build:
	go build -o bin/go-blockchain

run: build
	bin/go-blockchain

printchain: build
	bin/go-blockchain printchain