package main

import (
	"bufio"
	"container/list"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	lines := list.New()
	for {
		line, err := reader.ReadSlice('\n')
		if err == io.EOF && len(line) == 0 {
			break
		}
		//printSlice(line)
		lineCopy := append([]uint8(nil), line...)
		lines.PushBack(lineCopy)
	}
	//printList(lines)
	firstLine := lines.Front().Value.([]uint8)
	start := findStart(firstLine)

	counter := 0
	beamSplit(lines.Front().Next(), start, &counter)

	printList(lines)
	fmt.Println(counter)
}

func printList(list *list.List) {
	for e := list.Front(); e != nil; e = e.Next() {
		fmt.Println(string(e.Value.([]uint8)))
	}
}

func printSlice(slice []uint8) {
	fmt.Println(string(slice))
}

func findStart(line []uint8) int {
	start := 0
	for line[start] != 'S' {
		start++
	}
	if start == len(line) {
		panic("Not a start line")
	}
	return start
}

func beamSplit(element *list.Element, col int, counter *int) {
	if element == nil {
		return // Bottom reached, no further lines to check for
	}
	line := element.Value.([]byte)
	prevLine := element.Prev().Value.([]uint8)
	if prevLine[col] == '|' {
		return // If lines already checked do not continue
	}
	prevLine[col] = '|'
	if line[col] == '^' {
		*counter++
		if col > 0 {
			beamSplit(element.Next(), col-1, counter)
		}
		if col < len(line)-1 {
			beamSplit(element.Next(), col+1, counter)
		}
		return
	}
	beamSplit(element.Next(), col, counter)
	return
}
