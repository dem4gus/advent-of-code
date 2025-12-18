package main

import (
	"log"
	"os"

	day01 "github.com/dem4gus/advent-of-code/2025/pkg"
)

func main() {
	errLog := log.New(os.Stderr, "error: ", log.LstdFlags|log.Lmsgprefix|log.Llongfile)

	day01.ProblemOne(errLog)
	day01.ProblemTwo(errLog)
}
