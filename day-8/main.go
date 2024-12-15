package main

import (
	"bufio"
	"log"
	"os"
	"slices"
	"time"
)

type Node struct {
	y int
	x int
}

func main() {
	start := time.Now()

	matrix := getInput()
	_, nodesMap := getNodes(matrix)

	allAntiNodes := make([]Node, 0)
	for _, nodes := range nodesMap {
		antiNodes := getAntiNodes(nodes, len(matrix[0]), len(matrix))
		for _, antiNode := range antiNodes {
			if !slices.Contains(allAntiNodes, antiNode) {
				allAntiNodes = append(allAntiNodes, antiNode)
			}
		}
	}
	log.Printf("Total anti-nodes: %d", len(allAntiNodes))
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

func getNodes(matrix []string) ([]string, map[string][]Node) {
	var nodeCharacters []string
	nodes := make(map[string][]Node)

	for i, line := range matrix {
		for j, char := range line {
			s := string(char)
			if s != "." {
				if !slices.Contains(nodeCharacters, s) {
					nodeCharacters = append(nodeCharacters, s)
				}
				nodes[s] = append(nodes[s], Node{x: j, y: i})
			}
		}
	}
	return nodeCharacters, nodes
}

func getAntiNodes(nodes []Node, maxX int, maxY int) []Node {
	antiNodes := []Node{}
	for i, node := range nodes {
		for j, node2 := range nodes {
			if i != j {
				xDiff := node2.x - node.x
				yDiff := node2.y - node.y
				antiNode1 := Node{x: node.x - xDiff, y: node.y - yDiff}
				if antiNode1.x >= 0 && antiNode1.x < maxX && antiNode1.y >= 0 && antiNode1.y < maxY {
					antiNodes = append(antiNodes, antiNode1)
				}
				antiNode2 := Node{x: node2.x + xDiff, y: node2.y + yDiff}
				if antiNode2.x >= 0 && antiNode2.x < maxX && antiNode2.y >= 0 && antiNode2.y < maxY {
					antiNodes = append(antiNodes, antiNode2)
				}
			}
		}
	}
	return antiNodes
}
