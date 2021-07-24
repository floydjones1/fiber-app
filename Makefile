
up:
	docker compose up -d

down:
	docker compose down

start:
	modd

tools:
	cd ~/. && go get github.com/cortesi/modd/cmd/modd