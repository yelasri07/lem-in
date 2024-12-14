package functions

import "fmt"

func (a *Info) Bfs(n string) {
	if a.Neiofstart == nil {
		a.Neiofstart = make(map[string]bool)
	}

	var queue [][]string

	queue = append(queue, []string{n})

	for len(queue) > 0 {

		path := queue[0]

		queue = queue[1:]

		lastroom := path[len(path)-1]

		if lastroom == a.End {

			newpath := append([]string{}, path...)
			a.UniquePaths = append(a.UniquePaths, newpath)
			a.FindTheBestPaths()
			break

		}

		for _, nei := range a.Tunnels[lastroom] {
			if !isvesited(path, nei) && nei != a.Start && ok(a.Res, nei) {
				newpath := append([]string{}, path...)
				newpath = append(newpath, nei)
				queue = append(queue, newpath)
			}
		}
	}
}

func (y *Info) FindTheBestPaths() {
	/// I will use this method to find the best paths to move the ants in it ....
	fmt.Println(y.NumberOfrooms)
	// nm ants . nm rooms 
}

func ok(n []string, a string) bool {
	for _, char := range n {
		if char == a {
			return false
		}
	}
	return true
}

func isvesited(path []string, room string) bool {
	for _, char := range path {
		if char == room {
			return true
		}
	}
	return false
}
