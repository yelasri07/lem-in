package functions

import (
	"sort"
	"strings"
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
			// if len(y.Res) == len(y.UniquePaths) {
			// 	S := y.UniquePaths
			// 	res := y.FindGroups(S)
			// 	y.UniquePaths = res
			// 	break
			// }
			
			if len(y.Res) == len(y.UniquePaths) {
				for _, p := range y.UniquePaths {
					y.FindGroups(p)
				}
				y.Gh()
				y.Kk()
				break
			}

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
func (y *Info) Kk() {
	
}

func (y *Info) Gh() {
	uniqueMap := make(map[string][]string)

	for _, row := range y.UniquePaths {
		h := row[1:]
		h = h[:len(h)-1]
		sortedRow := append([]string{}, row...)
		sort.Strings(sortedRow)

		key := strings.Join(sortedRow, ",")

		if existing, exists := uniqueMap[key]; !exists || len(row) < len(existing) {
			uniqueMap[key] = row
		}
	}

	result := [][]string{}
	for _, row := range uniqueMap {
		result = append(result, row)
	}

	y.UniquePaths = result
}

func (y *Info) FindGroups(FirstPath []string) {
	res := y.UniquePaths

	var ff [][]string

	ff = append(ff, FirstPath)

	for i := 0; i < len(res); i++ {
		Status := FindTheUniquePaths(FirstPath, res[i])
		if Status {
			ff = append(ff, res[i])
			continue
		} else if !Status {
			if len(y.Tunnels[res[i][0]]) > 2 {

				h := Bfs(y.Tunnels, y.End, y.Start, res[i][0])
				b := jj(h, res[i])
				// fmt.Println(b)
				ff = append(ff, b...)

			}
		}
	}

	// y.UniquePaths =append(y.UniquePaths, ff...)
	y.UniquePaths = ff
}

func jj(h [][]string, g []string) [][]string {
	var rr [][]string
	for _, p := range h {
		if FindTheUniquePaths(p, g) {
			rr = append(rr, p)
			rr = append(rr, p)
		} else {
			rr = append(rr, p)
		}
	}
	return rr
}

func FindTheUniquePaths(p1, p2 []string) bool {
	for i := 0; i < len(p1)-1; i++ {
		for j := 0; j < len(p2)-1; j++ {
			if p1[i] == p2[j] {
				return false
			}
		}
	}
	return true
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

func Bfs(g map[string][]string, End string, Start string, r string) [][]string {
	var res [][]string
	var queue [][]string

	queue = append(queue, []string{r})

	for len(queue) > 0 {

		path := queue[0]

		queue = queue[1:]

		lastroom := path[len(path)-1]

		if lastroom == End {

			newpath := append([]string{}, path...)
			res = append(res, newpath)
		}

		for _, nei := range g[lastroom] {
			if !isvesited(path, nei) && nei != r && nei != Start {
				newpath := append([]string{}, path...)
				newpath = append(newpath, nei)
				queue = append(queue, newpath)
			}
		}
	}
	return res
}
