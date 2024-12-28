package main

import (
	"bufio"
	"log"
	_ "net/http/pprof"
	"os"
	"runtime/pprof"
	"strconv"
	"strings"
	"time"
)

func main() {
	f, err := os.Create("cpu.prof")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	if err := pprof.StartCPUProfile(f); err != nil {
		log.Fatal(err)
	}
	defer pprof.StopCPUProfile()

	start := time.Now()

	ints := getInput()
	freqs := make(map[int]int)
	results := make(map[int][]int)
	for _, i := range ints {
		freqs[i] = 1
	}
	log.Printf("Input: %v", ints)
	for i := 0; i < 75; i++ {
		freqs, results = blink(freqs, results)
		log.Printf("Iteration: %v", i)
		sum := 0
		for _, v := range freqs {
			sum += v
		}
		log.Printf("Stones: %v", sum)
	}

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
			b = append(b, int(in))
		}
		return b
	}
	return nil
}

func blink(freqs map[int]int, results map[int][]int) (map[int]int, map[int][]int) {
	newFreqs := make(map[int]int)
	for k, v := range freqs {
		result, ok := results[k]
		if !ok {
			result = calc(k)
			results[k] = result
		}
		for _, r := range result {
			existingFreq, ok := newFreqs[r]
			if !ok {
				newFreqs[r] = v
			} else {
				newFreqs[r] = existingFreq + v
			}
		}
	}
	return newFreqs, results
}

func calc(v int) []int {
	if v == 0 {
		return []int{1}
	} else {
		l := intLength(v)
		if l%2 == 0 {
			halfLength := l / 2
			divisor := intPow(10, halfLength)
			firstHalf := v / divisor
			lastHalf := v % divisor
			return []int{firstHalf, lastHalf}
		} else {
			return []int{v * 2024}
		}
	}
}

func intLength(v int) int {
	length := 0
	temp := v
	for temp > 0 {
		length++
		temp /= 10
	}
	return length
}

func firstHalf(v int, length int) int {
	halfLength := length / 2
	divisor := intPow(10, halfLength)
	firstHalf := v / divisor

	return firstHalf
}

func lastHalf(v int, length int) int {
	halfLength := length / 2
	divisor := intPow(10, halfLength)
	lastHalf := v % divisor

	return lastHalf
}

func intPow(base int, exp int) int {
	result := 1
	for exp > 0 {
		if exp&1 == 1 {
			result *= base
		}
		base *= base
		exp >>= 1
	}
	return result
}
