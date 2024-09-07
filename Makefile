all: build test

dev:
	$$(go env GOPATH)/bin/air

build:
	go build -v -o main

test:
	go test -v .

clean:
	rm -f main

install:
	go install .

deploy:
	podman-compose build
	podman-compose up