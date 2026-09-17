package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	points := extractCoordinates("input")

	edges := buildEdges(points)
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].D < edges[j].D
	})

	ds := NewDSU(len(points))

	// Part 2

	i := 0
	for i < len(edges) {
		ds.Union(edges[i].A, edges[i].B)
		prevRoot := ds.Find(0)
		j := 0
		for j < len(points) {
			root := ds.Find(j)
			if root != prevRoot {
				break
			}
			j++
		}
		if j == len(points) {
			break
		}
		i++
	}

	if i != len(edges) {
		fmt.Println(points[edges[i].A].X * points[edges[i].B].X)
	} else {
		fmt.Println("There is no possibilities for a single tree")
	}

	// Part 1

	//maxEdges := 1000
	//for i := 0; i < maxEdges && i < len(edges); i++ {
	//	ds.Union(edges[i].A, edges[i].B)
	//}

	//// Build final circuits
	//circuits := make(map[int]int)
	//for i := range points {
	//	root := ds.Find(i)
	//	circuits[root]++
	//}
	//
	//fmt.Println("Circuit sizes:", circuits)
	//
	//var sizes []int
	//for _, size := range circuits {
	//	sizes = append(sizes, size)
	//}
	//
	//// Get the three largest circuits
	//top3 := topK(sizes, 3)
	//
	//fmt.Println("Three largest circuits:", top3)
	//fmt.Println("Product:", top3[0]*top3[1]*top3[2])
}

func topK(values []int, k int) []int {
	sort.Slice(values, func(i, j int) bool {
		return values[i] > values[j] // descending
	})
	if len(values) < k {
		return values
	}
	return values[:k]
}

// Coordinates structures

type Coordinates struct {
	X, Y, Z int64
}

func (c Coordinates) String() string {
	return fmt.Sprintf("(%d, %d, %d)", c.X, c.Y, c.Z)
}

func (c Coordinates) Distance(o Coordinates) int64 {
	dx := c.X - o.X
	dy := c.Y - o.Y
	dz := c.Z - o.Z
	return dx*dx + dy*dy + dz*dz
}

// DSU
type DSU struct {
	parent []int
	rank   []int
}

func NewDSU(n int) *DSU {
	p := make([]int, n)
	r := make([]int, n)
	for i := range p {
		p[i] = i
	}
	return &DSU{parent: p, rank: r}
}

func (d *DSU) Find(x int) int {
	if d.parent[x] != x {
		d.parent[x] = d.Find(d.parent[x])
	}
	return d.parent[x]
}

func (d *DSU) Union(a, b int) {
	ra := d.Find(a)
	rb := d.Find(b)

	if ra == rb {
		return
	}

	if d.rank[ra] < d.rank[rb] {
		d.parent[ra] = rb
	} else if d.rank[ra] > d.rank[rb] {
		d.parent[rb] = ra
	} else {
		d.parent[rb] = ra
		d.rank[ra]++
	}
}

// Edges

type Edge struct {
	A, B int
	D    int64
}

func buildEdges(points []Coordinates) []Edge {
	var edges []Edge
	for i := range points {
		for j := i + 1; j < len(points); j++ {
			edges = append(edges, Edge{
				A: i,
				B: j,
				D: points[i].Distance(points[j]),
			})
		}
	}
	return edges
}

// Read file

func extractCoordinates(input string) []Coordinates {
	file, err := os.Open(input)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	var boxes []Coordinates
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
		x, _ := strconv.ParseInt(coordinatesSlice[0], 10, 64)
		y, _ := strconv.ParseInt(coordinatesSlice[1], 10, 64)
		z, _ := strconv.ParseInt(coordinatesSlice[2], 10, 64)
		boxes = append(boxes, Coordinates{X: x, Y: y, Z: z})
	}
	return boxes
}
