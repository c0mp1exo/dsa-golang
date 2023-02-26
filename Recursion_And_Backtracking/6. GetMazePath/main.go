package main

import "fmt"

func main() {
	res := getMazePath(0, 0, 3, 3)
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
