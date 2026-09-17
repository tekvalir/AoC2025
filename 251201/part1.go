package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	lock := 50
	clicks := 0
	for scanner.Scan() {
		instructions := scanner.Text()
		i, err := strconv.Atoi(instructions[1:])
		if err != nil {
			panic(err)
		}
		switch instructions[0] {
		case 'L':
			lock -= i
		case 'R':
			lock += i
		}
		lock %= 100
		if lock == 0 {
			clicks++
		}
	}
	fmt.Println(clicks)
}
