package main

import "fmt"

func main() {
	printPermutations("abc", "")
}

func printPermutations(ros, ans string) {
	if len(ros) == 0 {
		fmt.Printf("%s\n", ans)
		return
	}

	for i := 0; i < len(ros); i++ {
		ch := ros[i]
		fh := ros[0:i]
		sh := ros[i+1:]
		roq := fh + sh
		printPermutations(roq, ans+string(ch))
	}
}
