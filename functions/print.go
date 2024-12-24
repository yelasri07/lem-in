package functions

import (
	"fmt"
	"strconv"
)

func (y *Info) Print(res [][]string) {
	res = Sort(res)
	NubOfAnts := conv(y.NumberOfAnts)

	// index := 0
	fmt.Println(res)
	for len(NubOfAnts) > 0 {
		for i := 0; i < len(res); i++ {
			if len(res) == 1 {
				res[i] = append(res[i], NubOfAnts[i])
				NubOfAnts = NubOfAnts[1:]
				continue
			}
			if i != len(res)-1 && len(res[i]) <= len(res[i+1]) {

				res[i] = append(res[i], NubOfAnts[i])
				NubOfAnts = NubOfAnts[1:]

			} else if i != len(res)-1 && len(res[i]) > len(res[i+1]) {

				res[i+1] = append(res[i+1], NubOfAnts[0])

				NubOfAnts = NubOfAnts[1:]

			}
		}
	}
	y.hhhhhhhh(res)
}

func (y *Info) hhhhhhhh(res [][]string) {
	for _, v := range res {
		fmt.Println(v)
	}
}

func conv(nm int) []string {
	var res []string

	for i := 1; i <= nm; i++ {
		l := strconv.Itoa(i)
		res = append(res, "L"+l)

	}
	return res
}

func Sort(res [][]string) [][]string {
	if len(res) == 1 {
		return res
	}
	for i := 0; i < len(res); i++ {
		if i != len(res)-1 && len(res[i]) > len(res[i+1]) {
			res[i], res[i+1] = res[i+1], res[i]
		} else {
			continue
		}
	}
	return res
}
