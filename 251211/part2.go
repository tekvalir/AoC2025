package main

import "fmt"

func partTwo() {
	path := PathFile("input")
	ch1, err1 := path.getLineChannel()
	if err1 != nil {
		panic(err1)
	}
	ids, neighbors := extractServers(ch1)
	fmt.Println(ids)
	checkIds := []int{ids["fft"], ids["dac"]}
	//paths := nbPathsWith(ids["svr"], ids["out"], neighbors, checkIds)
	fmt.Println(checkIds)
	fmt.Println(neighbors)
}
