package main

import (
	"fmt"
	"lee/aoc/helpers"
	"os"
	"sort"
	"strconv"
	"strings"
)

// s/o to @megagator
const ROUNDS = 10_000

type Monkey struct {
	items       helpers.Stack[int]
	operator    string
	amount      string
	divisor     int
	targets     map[bool]int
	inspections int
}

func (m *Monkey) Operate(index int, mod_len int) {
	amt, _ := strconv.Atoi(m.amount)
	if m.amount == "old" {
		amt = m.items[index]
	}
	switch m.operator {
	case "*":
		m.items[index] *= amt
	case "+":
		m.items[index] += amt
	}

	m.items[index] = m.items[index] % mod_len
}

func (m *Monkey) Inspect(ms []*Monkey, mod_len int) {
	for i := range m.items {
		m.Operate(i, mod_len)
		m.inspections++
	}

	for _, item := range m.items {
		throw_to := m.targets[item%m.divisor == 0]
		ms[throw_to].items.Push(m.items.Shift())
	}
}

func find_solution(content string) int {
	instruction := []string{}
	var curr_monkey *Monkey
	monkeys := []*Monkey{}
	inspections := []int{}
	item := 0
	product_of_divisors := 1

	for _, monkey_info := range strings.Split(content, "\n\n") {
		curr_monkey = &Monkey{targets: map[bool]int{}}
		monkeys = append(monkeys, curr_monkey)

		for step, raw_info := range strings.Split(monkey_info, "\n") {

			switch step {
			case 1:
				instruction = strings.Split(raw_info, ": ")
				for _, raw_item := range strings.Split(instruction[1], ", ") {
					item, _ = strconv.Atoi(raw_item)
					curr_monkey.items.Push(item)
				}
			case 2:
				instruction = strings.Split(raw_info, "Operation: new = old ")
				instruction = strings.Split(instruction[1], " ")
				curr_monkey.operator = instruction[0]
				curr_monkey.amount = instruction[1]
			case 3:
				instruction = strings.Split(raw_info, "Test: divisible by ")
				curr_monkey.divisor, _ = strconv.Atoi(instruction[1])
				product_of_divisors *= curr_monkey.divisor
			case 4:
				instruction = strings.Split(raw_info, "If true: throw to monkey ")
				item, _ = strconv.Atoi(instruction[1])
				curr_monkey.targets[true] = item
			case 5:
				instruction = strings.Split(raw_info, "If false: throw to monkey ")
				item, _ = strconv.Atoi(instruction[1])
				curr_monkey.targets[false] = item
			}
		}
	}

	for i := 0; i < ROUNDS; i++ {
		for _, monkey := range monkeys {
			monkey.Inspect(monkeys, product_of_divisors)
		}
	}

	for _, monkey := range monkeys {
		inspections = append(inspections, monkey.inspections)
	}

	sort.Slice(inspections, func(i, j int) bool {
		return inspections[i] > inspections[j]
	})

	return inspections[0] * inspections[1]
}

func main() {
	// read input
	content, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Print(fmt.Errorf("%w", err))
	}
	content_str := string(content)

	fmt.Println("Level of monkey business: ", find_solution(content_str))
}
