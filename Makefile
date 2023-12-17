.PHONY: run

run: build
	./bin/urunan

run-fresh: build
	./bin/urunan refreshdb

build:
	go build -o ./bin/ .