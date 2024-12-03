package main

import (
	"log"
	"os"
	"regexp"
	"strconv"
	"time"
)

func main() {
	start := time.Now()

	duration := time.Since(start)
	str := getInput()
	result := getMults(str)
	log.Printf("Result: %d\n", result)
	log.Printf("Execution time: %s\n", duration)
}

func getInput() string {
	b, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
	str := string(b)

	return str
}

func getMults(s string) int {
	r := regexp.MustCompile(`mul\(\d+,\d+\)`)
	r2 := regexp.MustCompile(`\d+\d*`)
	matches := r.FindAllString(s, -1)
	totalResult := 0
	for _, match := range matches {
		matches2 := r2.FindAllString(match, -1)
		int1, _ := strconv.Atoi(matches2[0])
		int2, _ := strconv.Atoi(matches2[1])
		result := int1 * int2
		totalResult += result
	}
	return totalResult
}
