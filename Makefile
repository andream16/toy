.PHONY: build test vendor run stop

build: test vendor
	docker build .

run: test vendor
	docker-compose up -d

stop: 
	docker-compose down --rmi local

vendor:
	go mod vendor

test:
	go fmt ./...
	go test -vet all ./... -count=1 -race
