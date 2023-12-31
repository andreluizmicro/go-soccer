docker-up:
	docker-compose up -d

docker-build:
	docker-compose build

docker-down: 
	docker-compose down

docker-bash:
	docker-compose exec app bash

go-test:
	go test ./... 

test-coverage:
	go test -coverprofile=./coverage/coverage.out ./internal/shared/domain/entity

show-coverage:
	go tool  cover -html=./coverage/coverage.out