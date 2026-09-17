package main

import (
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
	"sync"
)

type point struct {
	x, y int
}

func (p point) areaUntil(other point) int {
	dx := other.x - p.x
	if dx < 0 {
		dx = -dx
	}
	dy := other.y - p.y
	if dy < 0 {
		dy = -dy
	}

	return (dx + 1) * (dy + 1)
}

func parsePoint(line string) point {
	parts := strings.Split(line, ",")
	x, _ := strconv.Atoi(parts[0])
	y, _ := strconv.Atoi(parts[1])
	return point{x, y}
}

func getPointChannel(input chan string) chan point {
	ch := make(chan point, cap(input))

	go func() {
		for line := range input {
			ch <- parsePoint(line)
		}
		close(ch)
	}()

	return ch
}

func stepOne(input chan string) (int, error) {
	points := []point{}
	maxArea := math.MinInt

	for p := range getPointChannel(input) {
		points = append(points, p)

		if len(points) == 1 {
			continue
		}

		areas := make([]int, len(points))
		for idx, p2 := range points {
			areas[idx] = p.areaUntil(p2)
		}

		max := slices.Max(areas)
		if max > maxArea {
			maxArea = max
		}
	}

	return maxArea, nil
}

type color int8

const (
	white color = iota
	green
	red
)

type matrix struct {
	tiles  [][]color
	points []point
}

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func (m *matrix) squareContainsWhiteTiles(sq segment) bool {
	xMin := min(sq[0].x, sq[1].x)
	yMin := min(sq[0].y, sq[1].y)
	xMax := max(sq[0].x, sq[1].x)
	yMax := max(sq[0].y, sq[1].y)

	for y := yMin; y < yMax; y++ {
		if m.tiles[y][xMin] == white || m.tiles[y][xMax] == white {
			return true
		}
	}

	for x := xMin; x < xMax; x++ {
		if m.tiles[yMin][x] == white || m.tiles[yMax][x] == white {
			return true
		}
	}

	return false
}

func (m *matrix) addReds() {
	for _, p := range m.points {
		m.addRed(p)
	}
}

func newMatrix(points []point) (m matrix) {
	coords := getMatrixSize(points)
	rows := coords[1] + 1
	cols := coords[0] + 1
	array := make([][]color, rows)
	for idx := range array {
		array[idx] = make([]color, cols)
	}
	m = matrix{
		tiles:  array,
		points: points,
	}

	return m
}

func (m *matrix) addRed(p point) {
	m.tiles[p.y][p.x] = red
}

func (m *matrix) makeBorders() {
	for i := 0; i < len(m.points)-1; i++ {
		m.makeBorder(m.points[i], m.points[i+1])
	}
	m.makeBorder(m.points[len(m.points)-1], m.points[0])
}

func (m *matrix) makeBorder(p1 point, p2 point) {
	if p1.x == p2.x {
		m.makeVerticalBorder(p1.x, p1.y, p2.y)
	} else {
		m.makeHorizontalBorder(p1.x, p2.x, p1.y)
	}
}

func (m *matrix) makeHorizontalBorder(x1, x2, y int) {
	if x1 > x2 {
		tmp := x1
		x1 = x2
		x2 = tmp
	}

	for i := x1; i <= x2; i++ {
		if m.tiles[y][i] == white {
			m.tiles[y][i] = green
		}
	}
}

func (m *matrix) makeVerticalBorder(x, y1, y2 int) {
	if y1 > y2 {
		tmp := y1
		y1 = y2
		y2 = tmp
	}

	for i := y1; i <= y2; i++ {
		if m.tiles[i][x] == white {
			m.tiles[i][x] = green
		}
	}
}

func (m *matrix) makeGreens() {
	m.makeBorders()

	iterCount := len(m.tiles) / MagicThreadCount

	parallel(func(thrId int) {
		for i := range iterCount {
			row := m.tiles[iterCount*thrId+i]
			inside := false
			lastWasWhite := false

			for idx, cell := range row {
				switch cell {
				case white:
					if inside {
						row[idx] = green
					}
				default:
					if lastWasWhite {
						inside = !inside
						fmt.Println(inside)
					}
				}
				lastWasWhite = cell == white
			}
		}
	})
}

func getMatrixSize(points []point) [2]int {
	var x, y int
	for _, p := range points {
		if p.x > x {
			x = p.x
		}
		if p.y > y {
			y = p.y
		}
	}
	return [2]int{x, y}
}

func (m *matrix) String() string {
	iterCount := len(m.points) / MagicThreadCount
	lines := make([]string, len(m.tiles))

	parallel(func(thrId int) {
		for i := range iterCount {
			row := m.tiles[iterCount*thrId+i]
			runes := make([]rune, len(row))
			for idx, col := range row {
				switch col {
				case white:
					runes[idx] = '.'
				case green:
					runes[idx] = 'X'
				case red:
					runes[idx] = '#'
				}
			}
			lines[iterCount*thrId+i] = string(runes)
		}
	})

	return strings.Join(lines, "\n")
}

type segment [2]point

func newSegment(p1, p2 point) segment {
	seg := [2]point{p1, p2}
	return seg
}

func (m *matrix) allValidSquaresWith(p1 point) chan segment {
	ch := make(chan segment)

	go func() {
		for _, p2 := range m.points {
			seg := newSegment(p1, p2)
			if p2 != p1 {
				ch <- seg
			}
		}
		close(ch)
	}()

	return ch
}

// 495 == 55 * 9
// 495 == 99 * 5
// 495 == 165 * 3
var MagicThreadCount = 9

func parallel(f func(int)) {
	var wg sync.WaitGroup
	for thrId := range MagicThreadCount {
		wg.Go(func() {
			f(thrId)
		})
	}

	wg.Wait()
}

func stepTwo(input chan string) (int, error) {
	points := []point{}
	for p := range getPointChannel(input) {
		points = append(points, p)
	}
	m := newMatrix(points)
	m.addReds()
	m.makeGreens()

	iterCount := len(m.points) / MagicThreadCount
	boys := make([]int, MagicThreadCount)

	parallel(func(thrId int) {
		largestBoy := math.MinInt
		for i := range iterCount {
			p := m.points[iterCount*thrId+i]

			squareCh := m.allValidSquaresWith(p)
			for sq := range squareCh {
				area := sq[0].areaUntil(sq[1])
				if area > largestBoy && !m.squareContainsWhiteTiles(sq) {
					largestBoy = area
				}
			}
		}
		boys[thrId] = largestBoy
	})

	return slices.Max(boys), nil
}

func main() {
	input := FilePath("input")
	ch2, err2 := input.GetLineChannel()
	if err2 != nil {
		fmt.Errorf("step two error: %s\n", err2)
		return
	}

	val2, errStep2 := stepTwo(ch2)
	if errStep2 != nil {
		fmt.Errorf("step two error: %s\n", errStep2)
		return
	}

	fmt.Printf("Step two result: %d\n", val2)
}
