package main

import (
	"container/list"
	"fmt"
	"sync"
	"time"
)

const THREADS = 10

func partOne() {
	path := PathFile("input")
	ch1, err1 := path.getLineChannel()
	if err1 != nil {
		panic(err1)
	}
	machines := extractMachines(ch1)
	fmt.Println("MultiThread")
	t := time.Now()
	depth := calculateAllDepthToTargetIndicators(machines)
	fmt.Println("depth: ", depth)
	fmt.Println("time: ", time.Since(t))
	fmt.Println("SingleThread")
	t = time.Now()
	depth = calculateSingleThread(machines)
	fmt.Println("depth: ", depth)
	fmt.Println("time: ", time.Since(t))
}

func depthToTargetIndicators(machine Machine) int {
	queue := list.New()
	for i := range len(machine.buttons) {
		indicators := make([]bool, len(machine.targetIndicators))
		queue.PushBack(NodeIndicator{buttonId: i, indicators: indicators, depth: 1})
	}
	for queue.Front() != nil {
		nextNode := queue.Remove(queue.Front()).(NodeIndicator)
		indicators := pushButtonIndicators(nextNode.indicators, machine.buttons[nextNode.buttonId])
		if machine.isTargetIndicators(indicators) {
			return nextNode.depth
		}
		for i := range len(machine.buttons) {
			copyIndicators := make([]bool, len(indicators))
			copy(copyIndicators, indicators)
			queue.PushBack(NodeIndicator{buttonId: i, indicators: copyIndicators, depth: nextNode.depth + 1})
		}
	}
	panic("should not reach here")
}

func calculateSingleThread(machines []Machine) int {
	c := 0
	for _, machine := range machines {
		c += depthToTargetIndicators(machine)
	}
	return c
}

func calculateAllDepthToTargetIndicators(machines []Machine) int {
	c := 0
	wg := sync.WaitGroup{}
	m := sync.Mutex{}
	for _, machine := range machines {
		wg.Go(func() {
			worker(&m, machine, &c)
		})
	}
	wg.Wait()
	return c
}

func worker(m *sync.Mutex, machine Machine, count *int) {
	depth := depthToTargetIndicators(machine)
	m.Lock()
	*count += depth
	m.Unlock()
}
