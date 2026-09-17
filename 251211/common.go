package main

import (
	"bufio"
	"os"
	"strings"
)

type PathFile string

func (f PathFile) getLineChannel() (chan string, error) {
	file, err := os.Open(string(f))
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(file)

	c := make(chan string)
	go (func() {
		defer file.Close()

		for scanner.Scan() {
			c <- scanner.Text()
		}

		close(c)
	})()

	return c, nil
}

func parseServer(input string, ids map[string]int, nextId *int) (int, []int) {
	splitted := strings.Split(input, ": ")
	name := splitted[0]
	neighbors := strings.Split(splitted[1], " ")
	if ids[name] == 0 {
		ids[name] = *nextId
		*nextId++
	}
	neighborsIds := make([]int, len(neighbors))
	for i, name := range neighbors {
		if ids[name] == 0 {
			ids[name] = *nextId
			*nextId++
		}
		neighborsIds[i] = ids[name]
	}
	return ids[name], neighborsIds
}

func extractServers(input chan string) (map[string]int, map[int][]int) {
	var ids = make(map[string]int)
	var neighbors = make(map[int][]int)
	nextId := 1
	for server := range input {
		id, serverNeighbors := parseServer(server, ids, &nextId)
		neighbors[id] = serverNeighbors
	}
	return ids, neighbors
}
