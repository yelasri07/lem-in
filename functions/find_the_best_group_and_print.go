package functions

import "fmt"

func (y *Info) FindTheBestGroup() {
	fmt.Println(y.NumberOfAnts)
	fmt.Println(y.NumberOfGroups)
	var res []int

	for _, v := range y.AllGroups {
		b := hi(v)
		if uq(res , b) {
			y.UniqueGroups = append(y.UniqueGroups, v)
		}
		res = append(res, b)
	}
	fmt.Println(res)
	for _, v := range y.UniqueGroups {
		fmt.Println(v)
	}
}

func uq(res []int , b int) bool {
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
