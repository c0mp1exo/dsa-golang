package main

func main() {

}

func getMazePathWithJump(sr, sc, dr, dc int) []string {

	if sc > dc && sr > dr {
		return *new([]string)
	}
	if sc == dc && sr == dr {
		var res []string
		res = append(res, "")
		return res
	}

	getH1 := getMazePathWithJump(sr, sc+1, dr, dc)
	getH2 := getMazePathWithJump(sr, sc+2, dr, dc)
	getH3 := getMazePathWithJump(sr, sc+3, dr, dc)

	getV1 := getMazePathWithJump(sr+1, sc, dr, dc)
	getV2 := getMazePathWithJump(sr+2, sc, dr, dc)
	getV3 := getMazePathWithJump(sr+3, sc, dr, dc)

	var res []string
	for _, val := range getH1 {
		res = append(res, "h1"+val)
	}
	for _, val := range getH2 {
		res = append(res, "h2"+val)
	}
	for _, val := range getH3 {
		res = append(res, "h3"+val)
	}
	for _, val := range getV1 {
		res = append(res, "v1"+val)
	}
	for _, val := range getV2 {
		res = append(res, "v2"+val)
	}
	for _, val := range getV3 {
		res = append(res, "v3"+val)
	}
	return res

}
