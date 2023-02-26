package main

import "fmt"

func main() {
	ans := getStairsPath(3)
	for _, a := range ans {
		fmt.Printf("%s\n", a)
	}
}

func getStairsPath(source int) []string {

	if source < 0 {
		return *new([]string)
	} else if source == 0 {
		var res []string
		res = append(res, "")
		return res
	}

	p1 := getStairsPath(source - 1)
	p2 := getStairsPath(source - 2)
	p3 := getStairsPath(source - 3)
	var p []string
	for _, val := range p1 {
		p = append(p, val+"1")
	}
	for _, val := range p2 {
		p = append(p, val+"2")
	}
	for _, val := range p3 {
		p = append(p, val+"3")
	}
	return p
}
