package main

import "fmt"

func main() {
	//Base Number
	x := 2

	//Exponent
	y := 10

	res := PowerLogarithmic(x, y)
	fmt.Printf("%d raised to the power %d is %d", x, y, res)
}

// power(x,y int) will calculate x raised to the power y
func Power(x, y int) int {
	if y < 0 {
		return 0
	}
	if y == 0 {
		return 1
	}

	return x * Power(x, y-1)

}

// powerLogarithmic
func PowerLogarithmic(x, y int) int {
	if y < 0 {
		return 0
	}

	if y == 0 {
		return 1
	}

	pxb2 := PowerLogarithmic(x, y/2)
	res := pxb2 * pxb2

	if y%2 != 0 {
		res *= x
	}
	return res
}
