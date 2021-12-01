package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	input, err := readInput("input")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("P1: ", part1(*input))
	fmt.Println("P2: ", part2(*input))
}

func part1(depths []int) int {
	increaseCount := 0
	for index := range depths {
		depth := depths[index]
		if index != len(depths)-1 {
			nextDepth := depths[index+1]
			if depth < nextDepth {
				increaseCount += 1
			}
		}
	}
	return increaseCount
}

func part2(depths []int) int {
	increaseCount := 0
	for index := range depths {
		if index < len(depths)-3 {
			firstSlidingWindow := depths[index] + depths[index+1] + depths[index+2]
			secondSlidingWindow := depths[index+1] + depths[index+2] + depths[index+3]
			if firstSlidingWindow < secondSlidingWindow {
				increaseCount += 1
			}
		}
		break
	}

	return increaseCount
}

func readInput(inputFile string) (*[]int, error) {
	file, err := os.Open(inputFile)
	if err != nil {
		return nil, fmt.Errorf("Failed to open file for reading: %v", err)
	}
	var input []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		number, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, fmt.Errorf("Failed to convert str to int: %v", err)
		}
		input = append(input, number)
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("Failed inside scanner: %v", err)
	}

	return &input, nil
}
