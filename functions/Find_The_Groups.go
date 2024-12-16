package functions

import "fmt"


func (y * Info) Ll() {
	if y.Aa() {
		fmt.Println("yessssssssssss")
	} else {
		y.Bjjfs(y.UniquePaths[0][0])
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


func (y *Info) Bjjfs(n string) {
	var queue [][]string

	queue = append(queue, []string{n})

	for len(queue) > 0 {

		path := queue[0]

		queue = queue[1:]

		lastroom := path[len(path)-1]

		if lastroom == y.End {

			newpath := append([]string{}, path...)

			y.UniquePaths = append(y.UniquePaths, newpath)

		}

		for _, nei := range y.Tunnels[lastroom] {
			if !isvesited(path, nei) && nei != y.Start {
				newpath := append([]string{}, path...)
				newpath = append(newpath, nei)
				queue = append(queue, newpath)
			}
		}
	}
}
