package main

import (
	"bufio"
	"log"
	"os"
	"time"
)

func main() {
	start := time.Now()

	s := getInput()
	ints := getBlocks(s)
	// log.Printf("Blocks: %v", ints)
	duration := time.Since(start)
	compressed := compressBlocks(ints)
	// log.Printf("Compressed: %v", compressed)
	checksum := calcChecksum(compressed)
	log.Printf("Checksum: %d", checksum)
	log.Printf("Execution time: %s", duration)
}

func getInput() string {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
	defer file.Close()

	reader := bufio.NewScanner(file)
	for reader.Scan() {
		s := reader.Text()
		return s
	}
	return ""
}

func getBlocks(s string) []int {
	free := false
	result := []int{}
	counter := 0
	for _, r := range s {
		if free {
			for i := 0; i < int(r-'0'); i++ {
				result = append(result, -1)
			}
		} else {
			for i := 0; i < int(r-'0'); i++ {
				result = append(result, counter)
			}
			counter++
		}
		free = !free
	}
	return result
}

func compressBlocks(ints []int) []int {
	max := len(ints) - 1
	compressed := []int{}

	for i := range ints {
		if max <= i {
			break
		}
		if ints[i] != -1 {
			compressed = append(compressed, ints[i])
		} else {
			pos, id := getLastMovableBlock(&ints, max)
			compressed = append(compressed, id)
			max = pos - 1
		}
	}
	return compressed
}

func getLastMovableBlock(ints *[]int, max int) (int, int) {
	iDeref := (*ints)
	for i := max; i >= 0; i-- {
		if iDeref[i] != -1 {
			return i, iDeref[i]
		}
	}
	log.Fatalf("No movable block found")
	return -1, -1
}

func calcChecksum(ints []int) int {
	checksum := 0
	for i, id := range ints {
		checksum += i * id
	}
	return checksum
}
