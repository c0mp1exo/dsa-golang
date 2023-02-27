package main

import "fmt"

func main() {
	fmt.Printf("Enter start index of yours : ")
	var sr, sc int
	fmt.Scanf("%d%d\n", &sr, &sc)
	fmt.Printf("Enter start index of yours : ")
	var dr, dc int
	fmt.Scanf("%d%d\n", &dr, &dc)

	res := getMazePath(sr, sc, dr, dc)
	fmt.Printf("Possible ways for you are ...  ")
	for _, val := range res {
		fmt.Printf("%s\n", val)
	}
}

func getMazePath(srRow, srCol, dsRow, dsCol int) []string {

	if srRow > dsRow || srCol > dsCol {
		return *new([]string)
	}

	if srRow == dsRow && srCol == dsCol {
		var res []string
		res = append(res, "")
		return res
	}

	hString := getMazePath(srRow+1, srCol, dsRow, dsCol)
	vString := getMazePath(srRow, srCol+1, dsRow, dsCol)

	var res []string

	for _, val := range hString {
		res = append(res, "h"+val)
	}
	for _, val := range vString {
		res = append(res, "v"+val)
	}
	return res

}
