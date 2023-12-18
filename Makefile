.PHONY: all
all: target/day01 target/day02 target/day03 target/day04 target/day05 target/day06 target/day07 target/day08 target/day09 target/day10 target/day11 target/day12 target/day13 target/day14 target/day15 target/day16 target/day17 target/day18

target/day01: $(shell find cmd/day01 -name "*.go") $(shell find internal -name "*.go")
	go build -o ./target/day01 ./cmd/day01

target/day02: $(shell find cmd/day02 -name "*.go") $(shell find internal -name "*.go")
	go build -o ./target/day02 ./cmd/day02

target/day03: $(shell find cmd/day03 -name "*.go") $(shell find internal -name "*.go")
	go build -o ./target/day03 ./cmd/day03

target/day04: $(shell find cmd/day04 -name "*.go") $(shell find internal -name "*.go")
	go build -o ./target/day04 ./cmd/day04

target/day05: $(shell find cmd/day05 -name "*.go") $(shell find internal -name "*.go")
	go build -o ./target/day05 ./cmd/day05

target/day06: $(shell find cmd/day06 -name "*.go") $(shell find internal -name "*.go")
	go build -o ./target/day06 ./cmd/day06

target/day07: $(shell find cmd/day07 -name "*.go") $(shell find internal -name "*.go")
	go build -o ./target/day07 ./cmd/day07

target/day08: $(shell find cmd/day08 -name "*.go") $(shell find internal -name "*.go")
	go build -o ./target/day08 ./cmd/day08

target/day09: $(shell find cmd/day09 -name "*.go") $(shell find internal -name "*.go")
	go build -o ./target/day09 ./cmd/day09

target/day10: $(shell find cmd/day10 -name "*.go") $(shell find internal -name "*.go")
	go build -o ./target/day10 ./cmd/day10

target/day11: $(shell find cmd/day11 -name "*.go") $(shell find internal -name "*.go")
	go build -o ./target/day11 ./cmd/day11

target/day12: $(shell find cmd/day12 -name "*.go") $(shell find internal -name "*.go")
	go build -o ./target/day12 ./cmd/day12

target/day13: $(shell find cmd/day13 -name "*.go") $(shell find internal -name "*.go")
	go build -o ./target/day13 ./cmd/day13

target/day14: $(shell find cmd/day14 -name "*.go") $(shell find internal -name "*.go")
	go build -o ./target/day14 ./cmd/day14

target/day15: $(shell find cmd/day15 -name "*.go") $(shell find internal -name "*.go")
	go build -o ./target/day15 ./cmd/day15

target/day16: $(shell find cmd/day16 -name "*.go") $(shell find internal -name "*.go")
	go build -o ./target/day16 ./cmd/day16

target/day17: $(shell find cmd/day17 -name "*.go") $(shell find internal -name "*.go")
	go build -o ./target/day17 ./cmd/day17

target/day18: $(shell find cmd/day18 -name "*.go") $(shell find internal -name "*.go")
	go build -o ./target/day18 ./cmd/day18

.PHONY: test
test:
	go test ./...

.PHONY: clean
clean:
	rm -rf target