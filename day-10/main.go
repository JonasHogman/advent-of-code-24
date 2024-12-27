package main

import (
	"bufio"
	"log"
	"os"
	"slices"
	"strconv"
	"sync/atomic"
	"time"
)

func main() {
	start := time.Now()

	matrix := getInput()
	startTrails(matrix)
	duration := time.Since(start)
	log.Printf("Execution time: %s", duration)
}

func getInput() []string {
	var matrix []string
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
	defer file.Close()

	reader := bufio.NewScanner(file)
	for reader.Scan() {
		s := reader.Text()
		matrix = append(matrix, s)
	}
	return matrix
}

func startTrails(matrix []string) {
	globalCounter := new(int32)
	*globalCounter = 0
	trailHeads := map[int][][]int{}

	counter := 0

	for i, line := range matrix {
		for j, char := range line {
			if char == '0' {
				findTrails(matrix, []int{i, j}, 0, globalCounter, trailHeads, counter)
				counter++
			}
		}
	}

	trailHeadCounter := 0
	for _, s := range trailHeads {
		slices.SortFunc(s, slices.Compare)
		s = slices.CompactFunc(s, slices.Equal)
		trailHeadCounter += len(s)
	}

	log.Printf("Trail head counter: %d", trailHeadCounter)
}

func findTrails(matrix []string, startPos []int, startNum int, globalCounter *int32, trailHeads map[int][][]int, counter int) {
	if startNum == 9 {
		atomic.AddInt32(globalCounter, 1)
		trailHeads[counter] = append(trailHeads[counter], startPos)
		return
	}
	targetNum := startNum + 1
	if startPos[0] > 0 {
		up := matrix[startPos[0]-1][startPos[1]]
		upInt, _ := strconv.Atoi(string(up))
		if upInt == targetNum {
			findTrails(matrix, []int{startPos[0] - 1, startPos[1]}, targetNum, globalCounter, trailHeads, counter)
		}
	}
	if startPos[0] < len(matrix)-1 {
		down := matrix[startPos[0]+1][startPos[1]]
		downInt, _ := strconv.Atoi(string(down))
		if downInt == targetNum {
			findTrails(matrix, []int{startPos[0] + 1, startPos[1]}, targetNum, globalCounter, trailHeads, counter)
		}
	}
	if startPos[1] > 0 {
		right := matrix[startPos[0]][startPos[1]-1]
		rightInt, _ := strconv.Atoi(string(right))
		if rightInt == targetNum {
			findTrails(matrix, []int{startPos[0], startPos[1] - 1}, targetNum, globalCounter, trailHeads, counter)
		}
	}
	if startPos[1] < len(matrix[0])-1 {
		left := matrix[startPos[0]][startPos[1]+1]
		leftInt, _ := strconv.Atoi(string(left))
		if leftInt == targetNum {
			findTrails(matrix, []int{startPos[0], startPos[1] + 1}, targetNum, globalCounter, trailHeads, counter)
		}
	}
}
