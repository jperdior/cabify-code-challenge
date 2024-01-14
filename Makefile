
tests:
	go test -v ./...

start: build run

restart : stop start

build:
	docker-compose build

run:
	docker-compose up -d

stop:
	docker-compose down
