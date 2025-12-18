package main

import aoc2025 "github.com/dem4gus/advent-of-code/2025/pkg"

type daySolver interface {
	Run()
}

func main() {
	problems := []daySolver{
		&aoc2025.Day01{},
	}

	for _, d := range problems {
		d.Run()
	}
}
