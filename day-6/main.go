package main

import (
	"bufio"
	"log"
	"os"
	"slices"
	"strings"
	"time"
)

type Position struct {
	x       int
	y       int
	visited bool
}

func main() {
	start := time.Now()

	matrix, x, y := getInput()
	visitedPositions := getVisitedPositions(matrix, x, y)
	positionCount := countVisited(visitedPositions)
	log.Printf("Distinct visited positions: %d", positionCount)

	duration := time.Since(start)
	log.Printf("Execution time: %s", duration)
}

func getInput() ([]string, int, int) {
	var matrix []string
	var x, y int
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
	defer file.Close()

	reader := bufio.NewScanner(file)
	i := 0
	for reader.Scan() {
		s := reader.Text()
		if strings.Contains(s, "^") {
			// get index of ^
			x = strings.Index(s, "^")
			y = i
		}
		matrix = append(matrix, s)
		i++
	}
	return matrix, x, y
}

func getVisitedPositions(matrix []string, x int, y int) map[int][]int {
	// 0, 0 is top left
	moves := 0
	visitedPositions := make(map[int][]int)
	currentDirection := "up"
	for {
		_, ok := visitedPositions[y]
		if !ok {
			visitedPositions[y] = []int{x}
		} else if !slices.Contains(visitedPositions[y], x) {
			visitedPositions[y] = append(visitedPositions[y], x)
		}

		switch currentDirection {
		case "up":
			if y == 0 {
				return visitedPositions
			}
			checkAhead := matrix[y-1][x]
			switch checkAhead {
			case '#':
				currentDirection = "right"
			case '.':
				y = y - 1
				moves++
			}
		case "right":
			if x == len(matrix[y])-1 {
				return visitedPositions
			}
			checkAhead := matrix[y][x+1]
			switch checkAhead {
			case '#':
				currentDirection = "down"
			case '.':
				x = x + 1
				moves++
			}
		case "down":
			if y == len(matrix)-1 {
				return visitedPositions
			}
			checkAhead := matrix[y+1][x]
			switch checkAhead {
			case '#':
				currentDirection = "left"
			case '.':
				y = y + 1
				moves++
			}
		case "left":
			if x == 0 {
				return visitedPositions
			}
			checkAhead := matrix[y][x-1]
			switch checkAhead {
			case '#':
				currentDirection = "up"
			case '.':
				x = x - 1
				moves++
			}
		default:
			return visitedPositions
		}
	}
}

func countVisited(visitedPositions map[int][]int) int {
	count := 0
	for _, v := range visitedPositions {
		count += len(v)
	}
	return count
}
