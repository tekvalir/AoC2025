package main

import (
	"container/heap"
	"fmt"
	"sync"
	"time"
)

func partTwo() {
	path := PathFile("input")
	ch1, err1 := path.getLineChannel()
	if err1 != nil {
		panic(err1)
	}
	machines := extractMachines(ch1)
	fmt.Println("MultiThread")
	t := time.Now()
	depth := calculateAllDepthToTargetJoltages(machines)
	//depth := calculateSingleThreadDepth(machines)
	//depth := depthToTargetJoltages(machines[0])
	fmt.Println("depth: ", depth)
	fmt.Println("time: ", time.Since(t))
	//fmt.Println("SingleThread")
	//t = time.Now()
	//depth = calculateSingleThread(machines)
	//fmt.Println("depth: ", depth)
	//fmt.Println("time: ", time.Since(t))
}

func heuristic(target []int, initial []int, depth, buttonLen int) int {
	d := 0
	for i, val := range initial {
		dx := target[i] - val
		dx = max(dx, -dx)
		d += dx * dx
	}
	return d - buttonLen + depth
}

func depthToTargetJoltages(machine Machine) int {
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	for i, button := range machine.buttons {
		counters := make([]int, len(machine.joltageCounters))
		heap.Push(&pq, &Item{node: NodeJoltage{buttonId: i, counters: counters, depth: 1}, priority: heuristic(counters, machine.joltageCounters, 1, len(button))})
	}
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		nextNode := (*item).node
		//if nextNode.depth > 11 {
		//	panic("end")
		//}
		//fmt.Println("initial counters: ", nextNode.counters)
		//fmt.Println("button pressed: ", nextNode.buttonId)
		//fmt.Println(nextNode.depth)
		//fmt.Println(item.priority)
		//fmt.Println("counters: ", counters)
		for i, button := range machine.buttons {
			counters := make([]int, len(nextNode.counters))
			copy(counters, nextNode.counters)
			pushButtonJoltage(counters, button)
			if machine.isTargetCounters(counters) {
				return nextNode.depth
			}
			if machine.shouldElagate(counters) {
				//fmt.Println("elagated")
				continue
			}
			heap.Push(&pq, &Item{node: NodeJoltage{buttonId: i, counters: counters, depth: nextNode.depth + 1}, priority: heuristic(counters, machine.joltageCounters, nextNode.depth+1, len(button))})
			//queue.PushBack(NodeJoltage{buttonId: i, counters: counters, depth: nextNode.depth + 1})
		}
	}
	panic("should not reach here")
}

func calculateSingleThreadDepth(machines []Machine) int {
	c := 0
	for i, machine := range machines {
		c += depthToTargetJoltages(machine)
		fmt.Println(i)
	}
	return c
}

func calculateAllDepthToTargetJoltages(machines []Machine) int {
	c := 0
	wg := sync.WaitGroup{}
	m := sync.Mutex{}
	for i, machine := range machines {
		wg.Go(func() {
			worker2(&wg, &m, machine, &c)
			fmt.Println(i)
		})
	}
	wg.Wait()
	return c
}

func worker2(wg *sync.WaitGroup, m *sync.Mutex, machine Machine, count *int) {
	//defer (*wg).Done()
	depth := depthToTargetJoltages(machine)
	m.Lock()
	*count += depth
	m.Unlock()
}

type Item struct {
	node            NodeJoltage
	priority, index int
}

type PriorityQueue []*Item

func (pq *PriorityQueue) Len() int {
	return len(*pq)
}

func (pq *PriorityQueue) Less(i, j int) bool {
	// We want the lowest priority (smallest integer) as the highest priority
	return (*pq)[i].priority < (*pq)[j].priority
}

func (pq *PriorityQueue) Swap(i, j int) {
	(*pq)[i], (*pq)[j] = (*pq)[j], (*pq)[i]
	(*pq)[i].index = i
	(*pq)[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}
