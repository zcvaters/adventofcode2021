package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	inputs, err := readInput("input")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("P1", part1(inputs))
	fmt.Println("P2", part2(inputs))
}

func part1(inputs []string) int {
	var instruction string
	var horizontal, vertical, units int
	for index := range inputs {
		_, err := fmt.Sscanf(inputs[index], "%s %d", &instruction, &units)
		if err != nil {
			break
		}
		switch instruction {
		case "forward":
			horizontal += units
		case "down":
			vertical += units
		case "up":
			vertical -= units
		}

	}

	return vertical * horizontal
}

func part2(inputs []string) int {
	var instruction string
	var horizontal, vertical, aim, units int
	for index := range inputs {
		_, err := fmt.Sscanf(inputs[index], "%s %d", &instruction, &units)
		if err != nil {
			break
		}
		switch instruction {
		case "forward":
			horizontal += units
			vertical += aim * units
		case "down":
			aim += units
		case "up":
			aim -= units
		}

	}

	return vertical * horizontal
}

func readInput(inputFile string) ([]string, error) {
	file, err := os.Open(inputFile)
	if err != nil {
		return nil, fmt.Errorf("Failed to open file for reading: %v", err)
	}
	scanner := bufio.NewScanner(file)
	var input []string
	for scanner.Scan() {
		if err != nil {
			return nil, fmt.Errorf("Failed to convert str to int: %v", err)
		}
		input = append(input, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("Failed inside scanner: %v", err)
	}

	return input, nil
}
