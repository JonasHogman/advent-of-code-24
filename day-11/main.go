package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()

	ints := getInput()
	log.Printf("Input: %v", ints)
	for i := 0; i < 25; i++ {
		ints = blink(ints)
	}

	stones := len(ints)
	log.Printf("Stones: %d", stones)

	duration := time.Since(start)
	log.Printf("Execution time: %s", duration)
}

func getInput() []int {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
	defer file.Close()

	reader := bufio.NewScanner(file)
	for reader.Scan() {
		s := reader.Text()
		a := strings.Split(s, " ")
		b := []int{}
		for _, v := range a {
			in, _ := strconv.Atoi(v)
			b = append(b, in)
		}
		return b
	}
	return nil
}

func blink(ints []int) []int {
	newInts := []int{}
	for _, v := range ints {
		if v == 0 {
			newInts = append(newInts, 1)
		} else {
			s := strconv.Itoa(v)
			l := len(s)
			if l%2 == 0 {
				firstHalf, _ := strconv.Atoi(s[:l/2])
				lastHalf, _ := strconv.Atoi(s[l/2:])
				newInts = append(newInts, firstHalf)
				newInts = append(newInts, lastHalf)
			} else {
				newInts = append(newInts, v*2024)
			}
		}
	}
	return newInts
}
