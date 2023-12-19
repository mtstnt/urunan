.PHONY: run

run: build
	./bin/urunan

run-fresh: build
	./bin/urunan refreshdb

run-fe:
	cd web && pnpm run dev

build:
	go build -o ./bin/ .