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

	p1Res, err := part1(inputs)
	if err != nil {
		fmt.Println(err)
		return
	}

	p2Res, err := part2(inputs)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("P1", *p1Res)
	fmt.Println("P2", *p2Res)
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

func part1(inputs []string) (*int, error) {
	var instruction string
	var horizontal, vertical, units int
	for index := range inputs {
		_, err := fmt.Sscanf(inputs[index], "%s %d", &instruction, &units)
		if err != nil {
			return nil, fmt.Errorf("Failed in scanner: %v", err)
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
	depth := vertical * horizontal
	return &depth, nil
}

func part2(inputs []string) (*int, error) {
	var instruction string
	var horizontal, vertical, aim, units int
	for index := range inputs {
		_, err := fmt.Sscanf(inputs[index], "%s %d", &instruction, &units)
		if err != nil {
			return nil, fmt.Errorf("Failed in scanner: %v", err)
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
	depth := vertical * horizontal
	return &depth, nil
}
