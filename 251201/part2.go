package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	debug := false
	if len(os.Args) > 1 && os.Args[1] == "--debug" {
		debug = true
	}
	file, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	lock := 50
	clicks := 0
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF && len(line) == 0 {
			break
		}
		instructions := strings.TrimSpace(line)
		start := lock
		if debug {
			fmt.Printf("initial lock %d, clicks %d, instructions %s", start, clicks, instructions)
		}
		steps, err := strconv.Atoi(instructions[1:])
		if err != nil {
			panic(err)
		}
		switch instructions[0] {
		case 'L':
			lock -= steps
		case 'R':
			lock += steps
		}
		clicks += steps / 100
		lock = (lock%100 + 100) % 100
		if lock == 0 || (lock < start && instructions[0] == 'R') || (lock > start && instructions[0] == 'L' && start != 0) {
			clicks++
		}
		if debug {
			fmt.Printf(" lock %d, clicks %d\n", lock, clicks)
		}
	}
	fmt.Println(clicks)
}
