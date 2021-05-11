.PHONY: start
start:
	docker-compose up -d; docker-compose logs -f

.PHONY: test
test:
	go test -v -cover -race -count=1 ./...