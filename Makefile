.PHONY: run

run: build
	./bin/urunan

build:
	go build -o ./bin/ .