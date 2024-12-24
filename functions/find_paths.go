package functions

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

func (y *Info) FindMorePaths(p []string) {
	if len(p) > 3 {
		n := p[0]
		b := p[1]
		y.BBfs(n, b)
	}
}

func (y *Info) BBfs(n string, b string) {
	var queue [][]string

	queue = append(queue, []string{n})

	for len(queue) > 0 {

		path := queue[0]

		queue = queue[1:]

		lastroom := path[len(path)-1]

		if len(y.Tunnels[lastroom]) > 3 {
			y.Bfs(lastroom)
			continue
		}

		if lastroom == y.End {

			newpath := append([]string{}, path...)
			y.UniquePaths = append(y.UniquePaths, newpath)

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

func ok(n []string, a string) bool {
	if len(n) == 0 {
		return true
	}
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
