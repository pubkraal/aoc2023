.PHONY: all
all: target/day01

target/day01: $(shell find . -name "*.go")
	go build -o ./target/day01 ./cmd/day01

