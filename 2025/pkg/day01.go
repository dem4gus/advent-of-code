package aoc2025

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Day01 struct{}

type rotations []int

func (d *Day01) Run() {
	fmt.Println("Day One")
	fmt.Println("------------")

	if rot := readFile("inputs/day01.txt"); rot != nil {
		rot.stepOne()
		rot.stepTwo()
	}
}

func (rot *rotations) stepOne() {
	dialPos := 50
	var password int

	for _, ticks := range *rot {
		dialPos += ticks
		if dialPos%100 == 0 {
			password += 1
		}
	}

	fmt.Println("Part 1:", password)
}

func (rot *rotations) stepTwo() {
	dialPos := 50
	var password int

	for _, ticks := range *rot {
		// add one every time the dial passes 0
		password += abs(ticks / 100)

		delta := ticks % 100
		finalPos := dialPos + delta

		// check for passing or landing on 0 one more time
		if (dialPos != 0 && finalPos <= 0) || finalPos >= 100 {
			password += 1
		}
		dialPos = (finalPos + 100) % 100
	}

	fmt.Printf("Part 2: %d\n", password)
}

func readFile(file string) *rotations {
	e := log.New(os.Stderr, "error: ", log.LstdFlags|log.Lmsgprefix|log.Llongfile)
	var data rotations

	f, err := os.Open(file)
	if err != nil {
		e.Println(err)
		return nil
	}

	defer func(f *os.File) {
		if err := f.Close(); err != nil {
			e.Println(err)
		}
	}(f)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		dir := line[0]
		ticks, err := strconv.Atoi(line[1:])
		if err != nil {
			e.Println(err)
			return nil
		}
		if dir == 'L' {
			ticks *= -1
		}
		data = append(data, ticks)
	}

	return &data
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
