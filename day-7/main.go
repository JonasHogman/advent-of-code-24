package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type Row struct {
	values []int
	result int
}

func main() {
	start := time.Now()

	rows := getInput()
	result := tryCombinations(rows)
	log.Printf("Result: %d", result)

	duration := time.Since(start)
	log.Printf("Execution time: %s", duration)
}

func getInput() []Row {
	var rows []Row

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
	defer file.Close()

	reader := bufio.NewScanner(file)
	for reader.Scan() {
		s := reader.Text()
		sp := strings.Split(s, ":")
		result, _ := strconv.Atoi(sp[0])
		if err != nil {
			log.Fatalf("failed converting string to int: %s", err)
		}
		numbers := strings.Fields(sp[1])
		ints := make([]int, len(numbers))
		for i, n := range numbers {
			converted, err := strconv.Atoi(n)
			if err != nil {
				log.Fatalf("failed converting string to int: %s", err)
			}
			ints[i] = converted
		}

		rows = append(rows, Row{values: ints, result: result})

	}
	return rows
}

func tryCombinations(rows []Row) int {
	validSum := 0

	for _, r := range rows {
		combinations := getAllOperationCombinations(len(r.values) - 1)
		if isValidCombination(combinations, r.values, r.result) {
			validSum += r.result
		}
	}
	return validSum
}

func isValidCombination(combinations [][]bool, values []int, k int) bool {
	for combination := range combinations {
		valid := calculate(values, combinations[combination], k)
		if valid {
			return true
		}
	}
	return false
}

func getAllOperationCombinations(k int) [][]bool {
	combinations := make([][]bool, 0)
	for i := 0; i < 1<<k; i++ {
		combination := make([]bool, k)
		for j := 0; j < k; j++ {
			combination[j] = (i>>j)&1 == 1
		}
		combinations = append(combinations, combination)
	}
	return combinations
}

func calculate(numbers []int, operations []bool, result int) bool {
	total := numbers[0]
	for i, number := range numbers[1:] {
		switch operations[i] {
		case true:
			total += number
		case false:
			total *= number
		}
		if total > result { // Break early if total exceeds result we are looking for
			return false
		}
	}
	if total == result {
		return true
	} else {
		return false
	}
}
