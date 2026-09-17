package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type coordinates struct {
	x, y, z int
}

func (c coordinates) String() string {
	return fmt.Sprintf("(%d, %d, %d)", c.x, c.y, c.z)
}

func (c coordinates) distance(c2 coordinates) int {
	dx := c.x - c2.x
	dy := c.y - c2.y
	dz := c.z - c2.z
	return dx*dx + dy*dy + dz*dz
}

func main() {
	boxes := readInput("input")
	distances := computeDistanceSemiMatrix(boxes)
	fmt.Println(distances)
	circuits := make([]int, len(distances))
	id := 1
	junctions(distances, &circuits, &id, 1000, boxes)
	fmt.Println(circuits)
	circuitsLen := make([]int, id)
	for i := range circuits {
		circuitsLen[circuits[i]]++
	}
	fmt.Println(circuitsLen)
}

func junctions(distances [][]int, circuits *[]int, nextId *int, nb int, coos []coordinates) {
	c := 0
	junc := map[string]bool{}
	for c < nb {
		minValue := -1
		var min1, min2 int
		for i := range distances {
			for j := i + 1; j < len(distances); j++ {
				s := strconv.Itoa(i) + "," + strconv.Itoa(j)
				if junc[s] {
					continue
				}
				if minValue == -1 || distances[i][j] < minValue {
					minValue = distances[i][j]
					min1 = i
					min2 = j
				}
			}
		}
		s := strconv.Itoa(min1) + "," + strconv.Itoa(min2)
		junc[s] = true
		//fmt.Printf("%d,%d,%d\n", minValue, min1, min2)
		//fmt.Printf("%s, %s\n", coos[min1].String(), coos[min2].String())
		if (*circuits)[min1] != 0 && (*circuits)[min2] != 0 {
			exCircuit := (*circuits)[min2]
			for i := range *circuits {
				if (*circuits)[i] == exCircuit {
					(*circuits)[i] = (*circuits)[min1]
				}
			}
		} else if (*circuits)[min1] != 0 {
			(*circuits)[min2] = (*circuits)[min1]
		} else if (*circuits)[min2] != 0 {
			(*circuits)[min1] = (*circuits)[min2]
		} else {
			(*circuits)[min1] = *nextId
			(*circuits)[min2] = *nextId
			*nextId++
		}
		//fmt.Println(*circuits)
		c++
	}
}

func computeDistanceSemiMatrix(coordinates []coordinates) [][]int {
	matrix := make([][]int, len(coordinates))
	for i := range matrix {
		matrix[i] = make([]int, len(coordinates))
	}
	for i := range coordinates {
		for j := i + 1; j < len(coordinates); j++ {
			matrix[i][j] = coordinates[i].distance(coordinates[j])
		}
	}
	return matrix
}

func readInput(input string) []coordinates {
	file, err := os.Open(input)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	var boxes []coordinates
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF && len(line) == 0 {
			break
		}
		line = strings.TrimSpace(line)
		coordinatesSlice := strings.Split(line, ",")
		if len(coordinatesSlice) != 3 {
			panic("invalid input at line: " + line)
		}
		fmt.Println(coordinatesSlice)
		x, _ := strconv.Atoi(coordinatesSlice[0])
		y, _ := strconv.Atoi(coordinatesSlice[1])
		z, _ := strconv.Atoi(coordinatesSlice[2])
		boxes = append(boxes, coordinates{x: x, y: y, z: z})
	}
	return boxes
}
