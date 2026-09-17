package main

import "runtime"

func main() {
	runtime.GOMAXPROCS(10)
	partOne()
	partTwo()
}
