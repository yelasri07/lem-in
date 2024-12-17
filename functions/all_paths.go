package functions

import (
	"fmt"
)

func (y *Info) Bfs(n string) {
	var queue [][]string

	queue = append(queue, []string{n})

	for len(queue) > 0 {

		path := queue[0]

		queue = queue[1:]

		lastroom := path[len(path)-1]

		if lastroom == y.End {

			newpath := append([]string{}, path...)

			y.UniquePaths = append(y.UniquePaths, newpath)
			break
		}

		for _, nei := range y.Tunnels[lastroom] {
			if !isvesited(path, nei) && nei != y.Start && ok(y.Res, nei) {
				newpath := append([]string{}, path...)
				newpath = append(newpath, nei)
				queue = append(queue, newpath)
			}
		}
	}
}

func (y *Info) FindGroups() {
	status, ig1, ig2 := FindTheUniquePaths(y.UniquePaths[0], y.UniquePaths[len(y.UniquePaths)-1])
	if status {
	} else {
		if len(y.Tunnels[y.UniquePaths[len(y.UniquePaths)-1][ig1-1]]) > 2 {
			N := y.UniquePaths[len(y.UniquePaths)-1][ig1-1]
			b := y.UniquePaths[len(y.UniquePaths)-1][ig1]
			y.BBfs(N, b)
		} else if len(y.Tunnels[y.UniquePaths[0][ig2-1]]) > 2 {
			fmt.Println(y.UniquePaths[0][ig2-1])
			N := y.UniquePaths[0][ig2-1]
			b := y.UniquePaths[0][ig2]
			y.BBfs(N, b)
		}
	}
}

func (y *Info) BBfs(n string, b string) {
	var queue [][]string

	queue = append(queue, []string{n})

	for len(queue) > 0 {

		path := queue[0]

		queue = queue[1:]

		lastroom := path[len(path)-1]

		if lastroom == y.End {

			newpath := append([]string{}, path...)

			y.UniquePaths = append(y.UniquePaths, newpath)
			break

		}

		for _, nei := range y.Tunnels[lastroom] {
			if !isvesited(path, nei) && nei != y.Start && ok(y.Res, nei) && nei != b {
				newpath := append([]string{}, path...)
				newpath = append(newpath, nei)
				queue = append(queue, newpath)
			}
		}
	}
}

func FindTheUniquePaths(p1, p2 []string) (bool, int, int) {
	for i := 1; i < len(p1)-1; i++ {
		for j := 1; j < len(p2)-1; j++ {
			if p1[i] == p2[j] {
				return false, j, i
			}
		}
	}
	return true, 0, 0
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

// ///////////////////////////////////
func FindTheUniqueePaths(p1, p2 []string) bool {
	for i := 0; i < len(p1)-1; i++ {
		for j := 0; j < len(p2)-1; j++ {
			if i != j && p1[i] == p2[j] {
				return false
			}
		}
	}
	return true
}

func (y *Info) Jj() {


	addedPaths := make(map[string]bool)

	for j := 0; j < len(y.UniquePaths); j++ {
		for i := 0; i < len(y.UniquePaths); i++ {
			if j != i && FindTheUniqueePaths(y.UniquePaths[j], y.UniquePaths[i]) {
				sortedPath := append([]string{}, y.UniquePaths[i]...)


				hh := jj(sortedPath)

				if !addedPaths[hh] {
					y.TheBestpaths = append(y.TheBestpaths, y.UniquePaths[i])
					addedPaths[hh] = true
				}
			}
		}
	}
}

func jj(res []string) string {
	b := ""
	for _, v := range res {
		b += v
	}
	return b
}
