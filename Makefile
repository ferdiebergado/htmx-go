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
	podman build -t htmx-go_app:latest .
	podman-compose up