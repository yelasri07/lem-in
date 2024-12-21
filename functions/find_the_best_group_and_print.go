package functions

import "fmt"

func (y *Info) FindTheBestGroup() {
	fmt.Println(y.NumberOfAnts)
	fmt.Println(y.NumberOfGroups)
	var res []int
	for _, v := range y.All {
		res = append(res, hi(v))
	}
	fmt.Println(res)
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
