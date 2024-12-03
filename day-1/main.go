package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"slices"
	"sort"
	"time"
)

func main() {
	start := time.Now()

	s1, s2 := getInput()
	totalDistance := calculateDistances(s1, s2)
	log.Printf("Total distance: %d", totalDistance)
	similarityScore := calculateSimilarityScore(s1, s2)
	log.Printf("Similarity Score: %d", similarityScore)

	duration := time.Since(start)
	log.Printf("Execution time: %s", duration)
}

func getInput() ([]int, []int) {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
	defer file.Close()

	var s1 []int
	var s2 []int

	reader := bufio.NewReader(file)
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatalf("failed reading file: %s", err)
		} else {
			var int1 int
			var int2 int
			for _, b := range line[0:5] { // since we know the input is always in the same format we can hardcode the positions
				int1 = int1*10 + int(b-48) // we also know that we only need to deal with ASCII 0-9
			}
			s1 = append(s1, int1)

			for _, b := range line[8:13] {
				int2 = int2*10 + int(b-48)
			}
			s2 = append(s2, int2)
		}
	}
	sort.Ints(s1)
	sort.Ints(s2)
	return s1, s2
}

func calculateDistances(s1 []int, s2 []int) int {
	totalDistance := 0
	for i := 0; i < len(s1); i++ {
		if s1[i] > s2[i] {
			totalDistance = totalDistance + (s1[i] - s2[i])
		} else {
			totalDistance = totalDistance + (s2[i] - s1[i])
		}
	}
	return totalDistance
}

func calculateSimilarityScore(s1 []int, s2 []int) int {
	score := 0
	for i := 0; i < len(s1); i++ {
		occurrences := getOccurrenceCount(s1[i], s2)
		score = score + (s1[i] * occurrences)
	}
	return score
}

func getOccurrenceCount(target int, sortedSlice []int) int {
	count := 0
	index, found := slices.BinarySearch(sortedSlice, target)

	// if we found one occurrence, scan left and right until we find a non-target value
	if found {
		count++
		for {
			if sortedSlice[index-count] == target {
				count++
			} else {
				break
			}
		}
		for {
			if sortedSlice[index+count] == target {
				count++
			} else {
				break
			}
		}
	}
	return count
}
