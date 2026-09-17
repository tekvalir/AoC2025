package main

import (
	"bufio"
	"os"
	"strings"
)

type Input interface {
	GetLineChannel() (chan string, error)
}

type FilePath string
type InputText string

func (path FilePath) GetLineChannel() (chan string, error) {
	file, err := os.Open(string(path))
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)

	ch := make(chan string, 1024)
	go (func() {
		defer file.Close()

		for scanner.Scan() {
			ch <- scanner.Text()
		}

		close(ch)
	})()

	return ch, nil
}

func (inputText InputText) GetLineChannel() (chan string, error) {
	lines := strings.Split(string(inputText), "\n")
	ch := make(chan string, len(lines))

	go (func() {
		for _, line := range lines {
			ch <- line
		}

		close(ch)
	})()

	return ch, nil
}
