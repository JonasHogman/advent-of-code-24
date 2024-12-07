package main

import (
	"bufio"
	"bytes"
	"log"
	"os"
	"time"
)

// x = 88
// m = 77
// a = 65
// s = 83

func main() {
	start := time.Now()

	matrix := getInput()
	scanMatrix(matrix)

	duration := time.Since(start)
	log.Printf("Execution time: %s", duration)
}

func getInput() [][]byte {
	var matrix [][]byte

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
	defer file.Close()

	reader := bufio.NewScanner(file)
	for reader.Scan() {
		b := reader.Text()
		matrix = append(matrix, []byte(b))
	}
	return matrix
}

func scanMatrix(matrix [][]byte) {
	matches := 0
	mas := []byte{77, 65, 83}
	for i := 0; i < len(matrix); i++ {
		for j := range matrix[i] {
			if matrix[i][j] == 88 {
				matches += scanForXmas(matrix, i, j, &mas)
			}
		}
	}
	log.Printf("Total matches: %v", matches)
}

func scanForXmas(matrix [][]byte, i int, j int, mas *[]byte) int {
	matches := 0

	if len(matrix[i]) > j+3 {
		horizontalRight := []byte{matrix[i][j+1], matrix[i][j+2], matrix[i][j+3]}
		if bytes.Equal(*mas, horizontalRight) {
			matches++
		}
	}
	// Horizontal left
	if j >= 3 {
		horizontalLeft := []byte{matrix[i][j-1], matrix[i][j-2], matrix[i][j-3]}
		if bytes.Equal(*mas, horizontalLeft) {
			matches++
		}
	}
	// Vertical down
	if i < len(matrix)-3 {
		verticalDown := []byte{matrix[i+1][j], matrix[i+2][j], matrix[i+3][j]}
		if bytes.Equal(*mas, verticalDown) {
			matches++
		}
	}
	// Vertical up
	if i >= 3 {
		verticalUp := []byte{matrix[i-1][j], matrix[i-2][j], matrix[i-3][j]}
		if bytes.Equal(*mas, verticalUp) {
			matches++
		}
	}
	// Diagonal right down
	if len(matrix[i]) > j+3 && i < len(matrix)-3 {
		diagonalRightDown := []byte{matrix[i+1][j+1], matrix[i+2][j+2], matrix[i+3][j+3]}
		if bytes.Equal(*mas, diagonalRightDown) {
			matches++
		}
	}
	// Diagonal right up
	if j+3 < len(matrix[i]) && i >= 3 {
		diagonalRightUp := []byte{matrix[i-1][j+1], matrix[i-2][j+2], matrix[i-3][j+3]}
		if bytes.Equal(*mas, diagonalRightUp) {
			matches++
		}
	}
	// Diagonal left down
	if j > 2 && i < len(matrix)-3 {
		diagonalLeftDown := []byte{matrix[i+1][j-1], matrix[i+2][j-2], matrix[i+3][j-3]}
		if bytes.Equal(*mas, diagonalLeftDown) {
			matches++
		}
	}
	// Diagonal left up
	if j > 2 && i > 2 {
		diagonalLeftUp := []byte{matrix[i-1][j-1], matrix[i-2][j-2], matrix[i-3][j-3]}
		if bytes.Equal(*mas, diagonalLeftUp) {
			matches++
		}
	}
	return matches
}
