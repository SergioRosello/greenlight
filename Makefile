all: build
	
build:
	go build -o greenlight ./cmd/api/

run: build
	export $(xargs <.env)
	./greenlight -cors-trusted-origins="http://localhost:9000 http://localhost:9001"
