package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Instruction struct {
	Direction string
	Distance  int
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			lines = append(lines, line)
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	var instructions []Instruction
	for _, line := range lines {
		if len(line) < 2 {
			log.Fatal("invalid line:", line)
		}
		dir := string(line[0])
		dist, err := strconv.Atoi(line[1:])
		if err != nil {
			log.Fatal("invalid distance:", line)
		}
		instructions = append(instructions, Instruction{Direction: dir, Distance: dist})
	}

	position := 50
	zeroCount := 0

	for _, inst := range instructions {
		if inst.Direction == "R" || inst.Direction == "r" {
			position = (position + inst.Distance) % 100
		} else if inst.Direction == "L" || inst.Direction == "l" {
			position = (position - inst.Distance + 100*1000) % 100
		} else {
			log.Fatal("invalid direction:", inst.Direction)
		}

		if position == 0 {
			zeroCount++
		}
	}

	fmt.Println("Password =", zeroCount)
}
