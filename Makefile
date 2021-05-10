.PHONY: start
start:
	docker-compose up -d; docker-compose logs -f

.PHONY: init
init:
	go get github.com/swaggo/swag/cmd/swag
	swag init -g cmd/users/users.go

.PHONY: test
test:
	go test -v -cover -race -count=1 ./...