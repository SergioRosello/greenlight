all: build
	
build:
	go build -o greenlight ./cmd/api/

run: build
	./greenlight
