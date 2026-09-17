package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var (
	mainPattern      = regexp.MustCompile("\\[(.*?)] (\\(.*\\)) \\{(\\d+(?:,\\d+)*)}")
	indicatorPattern = regexp.MustCompile("[.#]")
	buttonPattern    = regexp.MustCompile("\\d+(?:,\\d+)*")
)

// Can be common to all

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

// Just common to this day
type Button []int

type Machine struct {
	targetIndicators []bool
	buttons          []Button
	joltageCounters  []int
}

func (m Machine) String() string {
	return fmt.Sprintf("TI: %v, B: %v, JC: %v", m.targetIndicators, m.buttons, m.joltageCounters)
}

func (m Machine) isTargetIndicators(indicators []bool) bool {
	for i, indicator := range m.targetIndicators {
		if indicator != indicators[i] {
			return false
		}
	}
	return true
}

func (m Machine) isTargetCounters(counters []int) bool {
	for i, counter := range m.joltageCounters {
		if counter != counters[i] {
			return false
		}
	}
	return true
}

func (m Machine) shouldElagate(counters []int) bool {
	for i, counter := range m.joltageCounters {
		if counter < counters[i] {
			return true
		}
	}
	return false
}

func parseMachine(line string) Machine {
	match := mainPattern.FindStringSubmatch(line)
	//fmt.Println(match)
	targetIndicatorsMatches := indicatorPattern.FindAllString(match[1], -1)
	buttonMatches := buttonPattern.FindAllString(match[2], -1)
	targetIndicators := make([]bool, len(targetIndicatorsMatches))
	for i, match := range targetIndicatorsMatches {
		if match == "#" {
			targetIndicators[i] = true
		} else {
			targetIndicators[i] = false
		}
	}
	buttons := make([]Button, len(buttonMatches))
	for i, match := range buttonMatches {
		actuators := strings.Split(match, ",")
		buttons[i] = make(Button, len(actuators))
		for j, actuator := range actuators {
			buttons[i][j], _ = strconv.Atoi(actuator)
		}
	}
	counters := strings.Split(match[3], ",")
	joltageCounters := make([]int, len(counters))
	for i, counter := range counters {
		joltageCounters[i], _ = strconv.Atoi(counter)
	}
	return Machine{targetIndicators: targetIndicators, buttons: buttons, joltageCounters: joltageCounters}
}

func getMachineChannel(input chan string) chan Machine {
	ch := make(chan Machine, cap(input))

	go func() {
		for line := range input {
			ch <- parseMachine(line)
		}
		close(ch)
	}()

	return ch
}

func extractMachines(input chan string) []Machine {
	var machines []Machine
	for machine := range getMachineChannel(input) {
		machines = append(machines, machine)
	}
	return machines
}

type NodeIndicator struct {
	buttonId   int
	indicators []bool
	depth      int
}

type NodeJoltage struct {
	buttonId int
	counters []int
	depth    int
}

func pushButtonIndicators(indicators []bool, button Button) []bool {
	for _, indicator := range button {
		indicators[indicator] = !indicators[indicator]
	}
	return indicators
}

func pushButtonJoltage(counters []int, button Button) []int {
	for _, counter := range button {
		counters[counter]++
	}
	return counters
}
