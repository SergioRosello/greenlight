all: build
	
build:
	go build -o greenlight ./cmd/api/

run: build
	export $(xargs <.env)
	./greenlight
