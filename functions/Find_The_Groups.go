package functions

import "fmt"


func (y * Info) Ll() {
	if y.Aa() {
		fmt.Println("yessssssssssss")
	}
}

func (y * Info) Aa() bool {
	return y.Nn() > y.NumberOfAnts
}

func (y * Info) Nn() int {
	roomss := 0
	for  _, p := range y.TheBestpaths {
		for i := 0; i < len(p)-1; i++ {
			roomss++
		}
	}
	fmt.Println(roomss)
	return roomss
}