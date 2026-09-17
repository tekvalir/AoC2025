package main

import "fmt"

func partOne() {
	path := PathFile("input")
	ch1, err1 := path.getLineChannel()
	if err1 != nil {
		panic(err1)
	}
	ids, neighbors := extractServers(ch1)
	fmt.Println(pathToOut(ids["you"], ids["out"], neighbors))
}

func pathToOut(idIn int, idOut int, neighbors map[int][]int) int {
	hist := map[int]int{}
	visited := map[int]bool{}
	return visit(neighbors, &hist, &visited, idIn, idOut)
}

func visit(neighbors map[int][]int, hist *map[int]int, visited *map[int]bool, id int, idOut int) int {
	if (*visited)[id] {
		return max((*hist)[id], 0)
	}
	(*visited)[id] = true
	if id == idOut {
		(*hist)[id] = 1
		return 1
	}
	c := 0
	for _, neighbor := range neighbors[id] {
		c += visit(neighbors, hist, visited, neighbor, idOut)
	}
	(*hist)[id] = c
	return c
}
