.PHONY: all
all: target/day01 target/day02 target/day03

target/day01: $(shell find cmd/day01 -name "*.go") $(shell find internal -name "*.go")
	go build -o ./target/day01 ./cmd/day01

target/day02: $(shell find cmd/day02 -name "*.go") $(shell find internal -name "*.go")
	go build -o ./target/day02 ./cmd/day02

target/day03: $(shell find cmd/day03 -name "*.go") $(shell find internal -name "*.go")
	go build -o ./target/day03 ./cmd/day03