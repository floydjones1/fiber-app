
up:
	docker compose up -d

down:
	docker compose down

start:
	modd

## make create-migration [name_of_migration_file] [sql/go]
create-migration:
	cd internal/migration && goose create $(filter-out $@,$(MAKECMDGOALS))

tools:
	cd ~/. && 
	go get github.com/cortesi/modd/cmd/modd
	go get github.com/pressly/goose/cmd/goose