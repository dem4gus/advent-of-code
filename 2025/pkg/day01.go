package day01

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func ProblemOne(l *log.Logger) {
	dialPos := 50
	var password int

	f, err := os.Open("inputs/day01.txt")
	if err != nil {
		l.Println(err)
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		dir := line[0]
		ticks, err := strconv.Atoi(line[1:])
		if err != nil {
			l.Println(err)
			return
		}
		delta := ticks % 100
		if dir == 'L' {
			delta *= -1
		}

		intermediatePos := dialPos + delta
		if intermediatePos%100 == 0 {
			password += 1
		}
		dialPos = (intermediatePos + 100) % 100
	}

	fmt.Printf("Part 1: %d\n", password)
}

func ProblemTwo(l *log.Logger) {
	dialPos := 50
	var password int

	f, err := os.Open("inputs/day01.txt")
	if err != nil {
		l.Println(err)
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	lineNumber := 0
	for scanner.Scan() {
		line := scanner.Text()
		lineNumber += 1
		dir := line[0]
		ticks, err := strconv.Atoi(line[1:])
		if err != nil {
			l.Println(err)
			return
		}
		// add one every time the dial passes 0
		password += ticks / 100

		delta := ticks % 100
		if dir == 'L' {
			delta *= -1
		}
		finalPos := dialPos + delta

		// check for passing or landing on 0 one more time
		if (dialPos != 0 && finalPos <= 0) || finalPos >= 100 {
			password += 1
		}
		dialPos = (finalPos + 100) % 100
	}

	fmt.Printf("Part 2: %d\n", password)
}
