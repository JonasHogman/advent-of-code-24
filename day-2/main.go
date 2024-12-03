package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()

	matrix := getInput()
	safeReports := getSafeReports(matrix)
	log.Printf("Safe reports: %d", safeReports)

	safeReportsWithProblemDampener := getSafeReportsWithProblemDampener(matrix)
	log.Printf("Safe reports with problem dampener: %d", safeReportsWithProblemDampener)

	duration := time.Since(start)
	log.Printf("Execution time: %s", duration)
}

func getInput() [][]int {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
	defer file.Close()

	var matrix [][]int

	reader := bufio.NewReader(file)
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatalf("failed reading file: %s", err)
		} else {
			splitString := strings.Split(string(line), " ")
			var row []int
			for _, s := range splitString {
				n, err := strconv.Atoi(s)
				if err != nil {
					log.Fatalf("failed converting string to int: %s", err)
				}
				row = append(row, n)
			}
			matrix = append(matrix, row)
		}
	}
	return matrix
}

func getSafeReports(matrix [][]int) int {
	safeReports := 0
	for i := 0; i < len(matrix); i++ {
		if isReportSafe(matrix[i]) {
			safeReports++
		}
	}
	return safeReports
}

func getSafeReportsWithProblemDampener(matrix [][]int) int {
	safeReports := 0
	for i := 0; i < len(matrix); i++ {
		report := matrix[i]
		for j := 0; j < len(report)-1; j++ {
			tempReport := append(report[:j], report[j+1:]...)
			if isReportSafe(tempReport) {
				safeReports++
				continue
			}
		}
	}
	return safeReports
}

func isReportSafe(report []int) bool {
	var increasing bool // true if increasing, false if decreasing
	var lastNumber int
	for i := 0; i < len(report); i++ {
		switch i {
		case 0:
			lastNumber = report[i]
		case 1:
			if report[i] > lastNumber && report[i]-lastNumber < 4 {
				increasing = true
				lastNumber = report[i]
			} else if report[i] < lastNumber && report[i]-lastNumber > -4 {
				increasing = false
				lastNumber = report[i]
			} else {
				return false
			}
		default:
			difference := report[i] - lastNumber
			if increasing && difference > 0 && difference < 4 {
				lastNumber = report[i]
			} else if !increasing && difference < 0 && difference > -4 {
				lastNumber = report[i]
			} else {
				return false
			}
		}
	}
	return true
}
