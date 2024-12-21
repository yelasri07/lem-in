package functions

import "fmt"

func (y *Info) FindTheBestGroup() {

	var res []int
	for _, v := range y.AllGroups {
		b := hi(v)
		if uq(res, b) {
			y.UniqueGroups = append(y.UniqueGroups, v)
			y.NumberOfGroups++
		}
		res = append(res, b)
	}

	for _, v := range y.UniqueGroups {
		status, intg := y.IsGoodGroup(v)

		if status {
			fmt.Println(intg)
		    y.Print(v, y.NumberOfAnts , y.End)
		
			break
		}

	}
}

func (y *Info) IsGoodGroup(matrix [][]string) (bool, int) {
	b := hi(matrix)
	return true, b
}

func uq(res []int, b int) bool {
	if res == nil {
		return true
	}
	for _, it := range res {
		if it == b {
			return false
		}
	}
	return true
}

func hi(g [][]string) int {
	c := 0
	for _, v := range g {
		for _, h := range v {
			if h != "" {
				c++
			}
		}
	}
	return c
}
