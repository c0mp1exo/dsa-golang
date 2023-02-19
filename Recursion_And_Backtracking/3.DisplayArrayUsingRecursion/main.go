package main

import "fmt"

func main() {
	arr := []int{5, 4, 3, 2, 1}
	// Display(arr, 0)
	DisplayReverse(arr, 0)
}

func Display(arr []int, idx int) {
	if idx >= len(arr) {
		return
	}
	fmt.Printf("Value at idx %d is %d\n", idx, arr[idx])
	Display(arr, idx+1)
}

func DisplayReverse(arr []int, idx int) {
	if idx >= len(arr) {
		return
	}
	DisplayReverse(arr, idx+1)
	fmt.Printf("Value at idx %d is %d\n", idx, arr[idx])
}
