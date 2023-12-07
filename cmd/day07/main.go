package main

import (
	"flag"
	"fmt"
	"os"
	"reflect"
	"slices"
	"sort"
	"strings"
	"time"

	"github.com/pubkraal/aoc2023/internal/force"
	"github.com/pubkraal/aoc2023/internal/input"
	"github.com/pubkraal/aoc2023/internal/util"
)

type hand struct {
	Cards string
	Bid   int
}

func main() {
	// Define flags
	flag.Parse()
	if flag.NArg() == 0 {
		flag.Usage()
		os.Exit(1)
	}

	filename := flag.Arg(0)
	input := force.Must(input.ReadFileToStringSlice(filename))

	hands := make([]hand, len(input))
	for idx, line := range input {
		hand := hand{}
		fmt.Sscanf(line, "%s %d", &hand.Cards, &hand.Bid)
		hands[idx] = hand
	}

	p1 := 0
	p2 := 0

	var t1 time.Duration
	var t2 time.Duration

	for _, part2 := range []bool{false, true} {
		st := time.Now()
		sort.Slice(hands, func(i, j int) bool {
			h1 := value(hands[i].Cards, part2)
			h2 := value(hands[j].Cards, part2)
			s1 := strength(h1, part2)
			s2 := strength(h2, part2)
			if s1 == s2 {
				return h1 < h2
			} else {
				return s1 < s2
			}
		})

		for idx, hand := range hands {
			if !part2 {
				p1 += hand.Bid * (idx + 1)
				t1 = time.Since(st)
			} else {
				p2 += hand.Bid * (idx + 1)
				t2 = time.Since(st)
			}
		}
	}

	fmt.Printf("01: %d (%s)\n", p1, t1)
	fmt.Printf("02: %d (%s)\n", p2, t2)
}

func strength(hand string, part2 bool) int {
	de := reflect.DeepEqual
	cards := make(map[rune]int)
	nums := make([]int, 0)
	for _, r := range hand {
		cards[r]++
	}

	if part2 {
		ks := util.GetKeys(cards)
		target := ks[0]
		for k, val := range cards {
			if k != '1' && (val > cards[target] || target == '1') {
				target = k
			}
		}
		if _, ok := cards['1']; ok {
			if target != '1' {
				cards[target] += cards['1']
				delete(cards, '1')
			}
		}
	}

	for _, v := range cards {
		nums = append(nums, v)
	}

	slices.Sort(nums)

	if de(nums, []int{5}) { // Johnny Five Aces
		return 10
	} else if de(nums, []int{1, 4}) { // Four of a kind
		return 9
	} else if de(nums, []int{2, 3}) { // Full house
		return 8
	} else if de(nums, []int{1, 1, 3}) { // Three of a kind
		return 7
	} else if de(nums, []int{1, 2, 2}) { // Two pairs
		return 6
	} else if de(nums, []int{1, 1, 1, 2}) { // One pair
		return 5
	} else if de(nums, []int{1, 1, 1, 1, 1}) { // High card
		return 4
	}

	return 0
}

func value(hand string, part2 bool) string {
	// A, K, Q, J, T
	repl := strings.Map(func(r rune) rune {
		switch r {
		case 'T':
			return '9' + 1
		case 'J':
			if !part2 {
				return '9' + 2
			} else {
				return '2' - 1
			}
		case 'Q':
			return '9' + 3
		case 'K':
			return '9' + 4
		case 'A':
			return '9' + 5
		default:
			return r
		}
	}, hand)
	return repl
}
