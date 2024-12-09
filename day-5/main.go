package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	rules, updates := getInput()

	validUpdates := handleUpdates(rules, updates)
	log.Printf("Valid updates: %d", len(validUpdates))

	total := addMiddleNums(validUpdates)
	log.Printf("Total: %d", total)

	duration := time.Since(start)
	log.Printf("Execution time: %s", duration)
}

func getInput() (map[int][]int, [][]int) {
	rules := make(map[int][]int)
	var updates [][]int
	var before int
	var after int

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
	defer file.Close()

	reader := bufio.NewScanner(file)
	for reader.Scan() {
		b := reader.Text()
		if strings.Contains(b, "|") {
			_, err := fmt.Sscanf(b, "%d|%d", &before, &after)
			if err != nil {
				log.Fatalf("failed scanning rule: %s", err)
			}
			rules[after] = append(rules[after], before)
		} else if b != "" {
			s := strings.Split(b, ",")
			a := make([]int, len(s))
			for i, v := range s {
				a[i], _ = strconv.Atoi(v)
			}
			updates = append(updates, a)
		}

	}
	return rules, updates
}

func handleUpdates(rules map[int][]int, updates [][]int) [][]int {
	validUpdates := make([][]int, 0)
	for _, update := range updates {
		valid := checkUpdateValidity(rules, update)
		if valid {
			validUpdates = append(validUpdates, update)
		}
		// break
	}
	return validUpdates
}

func checkUpdateValidity(rules map[int][]int, update []int) bool {
	for i, currentNumber := range update { // iterate over each value in the update
		previousValues := update[:i] // get all values before the current number
		rule, ok := rules[currentNumber]
		if !ok {
			continue // no rules for this number
		} else {
			for _, ruleValue := range rule {
				if slices.Contains(update, ruleValue) {
					// log.Printf("Rule value %d is in the update", ruleValue)
					if !slices.Contains(previousValues, ruleValue) {
						return false
					} // if the rule value is not in the previous values, the update is invalid
				}
			}
		}
	}
	return true
}

func addMiddleNums(updates [][]int) int {
	total := 0
	for _, update := range updates {
		total += update[len(update)/2]
	}
	return total
}
