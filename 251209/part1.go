package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Coordinates struct {
	X, Y int64
}

func (c Coordinates) String() string {
	return fmt.Sprintf("(%d, %d)", c.X, c.Y)
}

func area(c1, c2 Coordinates) int64 {
	w := max(c1.X-c2.X, c2.X-c1.X) + 1
	h := max(c1.Y-c2.Y, c2.Y-c1.Y) + 1
	return w * h
}

type Cases struct {
	A, B Coordinates
	area int64
}

func (c Cases) String() string {
	return fmt.Sprintf("(%v, %v, %v)", c.A, c.B, c.area)
}

func runPart1() {
	coordinates := extractCoordinates("input")
	areas := computeAllAreas(coordinates)
	//sort.Slice(areas, func(i, j int) bool {
	//	return areas[i].area > areas[j].area
	//})
	//print(areas[0].String())
	m := slices.MaxFunc(areas, func(a, b Cases) int {
		if a.area > b.area {
			return 1
		} else if a.area < b.area {
			return -1
		}
		return 0
	})
	fmt.Println(m.area)
}

func computeAllAreas(coordinates []Coordinates) []Cases {
	var areas []Cases
	for i := range coordinates {
		for j := i + 1; j < len(coordinates); j++ {
			areas = append(areas, Cases{
				A:    coordinates[i],
				B:    coordinates[j],
				area: area(coordinates[i], coordinates[j]),
			})
		}
	}
	return areas
}

func extractCoordinates(input string) []Coordinates {
	file, err := os.Open(input)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	var coos []Coordinates
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF && len(line) == 0 {
			break
		}
		line = strings.TrimSpace(line)
		coordinatesSlice := strings.Split(line, ",")
		if len(coordinatesSlice) != 2 {
			panic("invalid input at line: " + line)
		}
		x, _ := strconv.ParseInt(coordinatesSlice[0], 10, 64)
		y, _ := strconv.ParseInt(coordinatesSlice[1], 10, 64)
		coos = append(coos, Coordinates{X: x, Y: y})
	}
	return coos
}
