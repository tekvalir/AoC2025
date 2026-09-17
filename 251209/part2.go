package main

import (
	"fmt"
	"math/rand/v2"
	"sort"
)

func runPart2() {
	coordinates := extractCoordinates("input")
	areas := computeAllAreas(coordinates)
	sort.Slice(areas, func(i, j int) bool {
		return areas[i].area > areas[j].area
	})
	for i := 0; i < len(areas); i++ {
		c := areas[i]
		X1 := c.A.X
		Y1 := c.A.Y
		X2 := c.B.X
		Y2 := c.B.Y
		fmt.Println(c)
		first := checkLine(coordinates, X1, Y1, X1, Y2)
		second := checkLine(coordinates, X1, Y2, X2, Y2)
		third := checkLine(coordinates, X2, Y2, X2, Y1)
		fourth := checkLine(coordinates, X2, Y1, X1, Y1)
		//fmt.Printf("%v, %v, %v, %v\n", first, second, third, fourth)
		if first && second && third && fourth {
			//if checkLine(&canvas, X1, Y1, X1, Y2) && checkLine(&canvas, X1, Y2, X2, Y2) && checkLine(&canvas, X2, Y2, X2, Y1) && checkLine(&canvas, X2, Y1, X1, Y1) {
			fmt.Println(c.area)
			return
		}

	}
}

//func checkLine(canvas *[][]byte, x1, y1, x2, y2 int64) bool {
//	//fmt.Printf("x1: %v, x2: %v, y1: %v, y2: %v\n", x1, x2, y1, y2)
//	if x1 == x2 && y1 == y2 {
//		return true
//	}
//	if x1 == x2 {
//		return checkLineY(canvas, x1, min(y1, y2), max(y1, y2))
//	} else if y1 == y2 {
//		return checkLineX(canvas, min(x1, x2), max(x1, x2), y1)
//	}
//	panic("NOPE")
//}
//
//func checkLineX(canvas *[][]byte, x1, x2, y int64) bool {
//	for i := x1 + 1; i < x2; i++ {
//		if !isIn(canvas, i, y) {
//			//fmt.Printf("%v, %v\n", i, y)
//			return false
//		}
//	}
//	return true

func checkLine(coordinates []Coordinates, X1, Y1, X2, Y2 int64) bool {
	if X1 == X2 && Y1 == Y2 {
		return true
	}
	if X1 == X2 {
		return checkLineY(coordinates, true, X1, Y1, Y2)
	} else {
		return checkLineX(coordinates, true, X1, X2, Y1)
	}
}

func checkLineY(coordinates []Coordinates, random bool, X, Y1, Y2 int64) bool {
	for i := Y1 + 1; i < Y2; i++ {
		if rand.IntN(10) != 0 && random {
			continue
		}
		if !isIn(coordinates, X, i) {
			return false
		}
	}
	return true
}

func checkLineX(coordinates []Coordinates, random bool, X1, X2, Y int64) bool {
	for i := X1 + 1; i < X2; i++ {
		if rand.IntN(10) != 0 && random {
			continue
		}
		if !isIn(coordinates, i, Y) {
			return false
		}
	}
	return true
}

func isIn(points []Coordinates, x, y int64) bool {
	c := 0
	for i := 0; i < len(points)-1; i++ {
		p1, p2 := points[i], points[i+1]
		if y == p1.Y && y == p2.Y && (min(p1.X, p2.X) <= x) && (max(p1.X, p2.X) >= x) {
			return true
		}
		if x == p1.X && x == p2.X && (min(p1.Y, p2.Y) <= y) && (max(p1.Y, p2.Y) >= y) {
			return true
		}

		// inner area
		if p1.Y == p2.Y && p1.Y < y {
			// vertical ray
			if min(p1.X, p2.X) < x && x <= max(p1.X, p2.X) {
				c++
			}
		}
	}
	return (c % 2) == 1
}

//func runPart2() {
//	coordinates := extractCoordinates("input")
//	//fmt.Println(coordinates)
//	mX := slices.MinFunc(coordinates, func(a, b Coordinates) int {
//		return compareInt64(a.X, b.X)
//	}).X
//	MX := slices.MaxFunc(coordinates, func(a, b Coordinates) int {
//		return compareInt64(a.X, b.X)
//	}).X
//	mY := slices.MinFunc(coordinates, func(a, b Coordinates) int {
//		return compareInt64(a.Y, b.Y)
//	}).Y
//	MY := slices.MaxFunc(coordinates, func(a, b Coordinates) int {
//		return compareInt64(a.Y, b.Y)
//	}).Y
//	//fmt.Printf("mX: %v, mY: %v\n", mX, mY)
//	//fmt.Printf("MX: %v, MY: %v\n", MX, MY)
//	//fmt.Println(coordinates)
//	offsetCoordinates(&coordinates, mX, mY)
//	//fmt.Println(coordinates)
//	w := MX - mX + 1
//	h := MY - mY + 1
//	canvas := drawMap(coordinates, w, h)
//	//printCanvas(canvas)
//	areas := computeAllAreas(coordinates)
//	sort.Slice(areas, func(i, j int) bool {
//		return areas[i].area > areas[j].area
//	})
//	i := 0
//	for i < len(areas) {
//		c := areas[i]
//		X1 := c.A.X
//		Y1 := c.A.Y
//		X2 := c.B.X
//		Y2 := c.B.Y
//		fmt.Println(c)
//		first := checkLine(&canvas, X1, Y1, X1, Y2)
//		second := checkLine(&canvas, X1, Y2, X2, Y2)
//		third := checkLine(&canvas, X2, Y2, X2, Y1)
//		fourth := checkLine(&canvas, X2, Y1, X1, Y1)
//		//fmt.Printf("%v, %v, %v, %v\n", first, second, third, fourth)
//		if first && second && third && fourth {
//			//if checkLine(&canvas, X1, Y1, X1, Y2) && checkLine(&canvas, X1, Y2, X2, Y2) && checkLine(&canvas, X2, Y2, X2, Y1) && checkLine(&canvas, X2, Y1, X1, Y1) {
//			fmt.Println(c.area)
//			return
//		}
//		i++
//	}
//}

//func compareInt64(a, b int64) int {
//	if a > b {
//		return 1
//	} else if a < b {
//		return -1
//	}
//	return 0
//}
//
//func offsetCoordinates(coordinates *[]Coordinates, offX, offY int64) {
//	for i := range *coordinates {
//		(*coordinates)[i].X = (*coordinates)[i].X - offX
//		(*coordinates)[i].Y = (*coordinates)[i].Y - offY
//	}
//}
//
//func newCanvas(w, h int64) [][]byte {
//	canvas := make([][]byte, h)
//	for i := range canvas {
//		canvas[i] = make([]byte, w)
//		for j := range canvas[i] {
//			canvas[i][j] = '.'
//		}
//	}
//	return canvas
//}
//
//func printCanvas(canvas [][]byte) {
//	for i := range canvas {
//		fmt.Println(string(canvas[i]))
//	}
//}
//
//func isIn(canvas *[][]byte, x, y int64) bool {
//	if (*canvas)[y][x] != '.' {
//		return true
//	}
//	if rand.IntN(100) != 0 {
//		return true
//	}
//	return ray(canvas, x, y, 0, 1) && ray(canvas, x, y, 0, -1) && ray(canvas, x, y, 1, 0) && ray(canvas, x, y, -1, 0)
//}
//
//func ray(canvas *[][]byte, x int64, y int64, dirX, dirY int) bool {
//	if dirX == 0 {
//		for i := int(y); i < len(*canvas) && i >= 0; i += dirY {
//			if (*canvas)[i][x] != '.' {
//				return true
//			}
//		}
//		return false
//	} else {
//		for i := int(x); i < len((*canvas)[y]) && i >= 0; i += dirX {
//			if (*canvas)[y][i] != '.' {
//				return true
//			}
//		}
//		return false
//	}
//
//}
//
//func checkLine(canvas *[][]byte, x1, y1, x2, y2 int64) bool {
//	//fmt.Printf("x1: %v, x2: %v, y1: %v, y2: %v\n", x1, x2, y1, y2)
//	if x1 == x2 && y1 == y2 {
//		return true
//	}
//	if x1 == x2 {
//		return checkLineY(canvas, x1, min(y1, y2), max(y1, y2))
//	} else if y1 == y2 {
//		return checkLineX(canvas, min(x1, x2), max(x1, x2), y1)
//	}
//	panic("NOPE")
//}
//
//func checkLineX(canvas *[][]byte, x1, x2, y int64) bool {
//	for i := x1 + 1; i < x2; i++ {
//		if !isIn(canvas, i, y) {
//			//fmt.Printf("%v, %v\n", i, y)
//			return false
//		}
//	}
//	return true
//}
//
//func checkLineY(canvas *[][]byte, x, y1, y2 int64) bool {
//	for i := y1 + 1; i < y2; i++ {
//		if !isIn(canvas, x, i) {
//			//fmt.Printf("%v, %v\n", x, i)
//			return false
//		}
//	}
//	return true
//}
//
//func color(canvas *[][]byte, c Coordinates) {
//	l := list.New()
//	l.PushFront(c)
//	for l.Front() != nil {
//		e := l.Remove(l.Front()).(Coordinates)
//		X := e.X
//		Y := e.Y
//		if X == -1 || Y == -1 || X == int64(len((*canvas)[Y])) || Y == int64(len(*canvas)) {
//			panic("NOT AN INSIDE POINT")
//		}
//		if (*canvas)[Y][X] != '.' {
//			continue
//		}
//		(*canvas)[Y][X] = 'C'
//		l.PushFront(Coordinates{X: X + 1, Y: Y})
//		l.PushFront(Coordinates{X: X, Y: Y + 1})
//		l.PushFront(Coordinates{X: X, Y: Y - 1})
//		l.PushFront(Coordinates{X: X - 1, Y: Y})
//	}
//}
//
//func drawMap(coordinates []Coordinates, w, h int64) [][]byte {
//	canvas := newCanvas(w, h)
//	for i := range coordinates {
//		X := coordinates[i].X
//		Y := coordinates[i].Y
//		canvas[Y][X] = '#'
//		var X2, Y2 int64
//		//fmt.Println(i)
//		if i == len(coordinates)-1 {
//			X2 = coordinates[0].X
//			Y2 = coordinates[0].Y
//		} else {
//			X2 = coordinates[i+1].X
//			Y2 = coordinates[i+1].Y
//		}
//		drawLine(&canvas, X, Y, X2, Y2)
//	}
//	return canvas
//}
//
//func drawLine(canvas *[][]byte, x1, y1, x2, y2 int64) {
//	if x1 == x2 && y1 == y2 {
//		return
//	}
//	if x1 == x2 {
//		drawLineY(canvas, x1, min(y1, y2), max(y1, y2))
//	} else if y1 == y2 {
//		drawLineX(canvas, min(x1, x2), max(x1, x2), y1)
//	}
//}
//
//func drawLineY(canvas *[][]byte, X, Y1, Y2 int64) {
//	//fmt.Printf("X: %v, Y1: %v, Y2: %v\n", X, Y1, Y2)
//	for i := Y1 + 1; i < Y2; i++ {
//		(*canvas)[i][X] = 'X'
//	}
//}
//
//func drawLineX(canvas *[][]byte, X1, X2, Y int64) {
//	//fmt.Printf("X1: %v, X2: %v, Y: %v\n", X1, X2, Y)
//	for i := X1 + 1; i < X2; i++ {
//		(*canvas)[Y][i] = 'X'
//	}
//}
