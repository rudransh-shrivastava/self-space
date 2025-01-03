build:
	go build -o bin/self-space

run: build
	./bin/self-space

test:
	go test -v ./... -count=1