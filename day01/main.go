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
func part1(instructions []Instruction) int {
	position := 50
	count := 0

	for _, inst := range instructions {
		if inst.Direction == "R" {
			position = (position + inst.Distance) % 100
		} else {
			position = (position - inst.Distance + 100*1000) % 100
		}

		if position == 0 {
			count++
		}
	}

	return count
}
func part2(instructions []Instruction) int {
	position := 50
	count := 0

	for _, inst := range instructions {
		for i := 0; i < inst.Distance; i++ {
			if inst.Direction == "R" {
				position = (position + 1) % 100
			} else {
				position = (position - 1 + 100) % 100
			}

			if position == 0 {
				count++
			}
		}
	}

	return count
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
    fmt.Println("Part 1 =", part1(instructions))
	fmt.Println("Part 2 =", part2(instructions))


}

