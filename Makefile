
up:
	docker compose up -d

down:
	docker compose down

start:
	modd

build: 
	go build -mod vendor -o ./bin/server ./cmd/server/main.go

## make create-migration [name_of_migration_file] [sql/go]
create-migration:
	./bin/goose create $(filter-out $@,$(MAKECMDGOALS))

## informs user about migration status
db-status:
	./bin/goose -dir=./internal/data/migrations status

## runs all migrations against DB
db-up:
	./bin/goose -dir=./internal/data/migrations up

## runs all migrations against DB
db-down:
	./bin/goose -dir=./internal/data/migrations down

tools:
	go build -o ./bin/goose ./cmd/goose/main.go && \
	cd ~ && \
	go get github.com/cortesi/modd/cmd/modd \
	github.com/vektra/mockery/v2/.../

create-mocks:
	rm -rf ./internal/data/mocks && go generate ./...