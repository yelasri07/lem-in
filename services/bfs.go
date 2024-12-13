package services

import "fmt"

func (g *GraphData) Bfs() {
	start := g.Start
	visited := map[string]bool{
		start: true,
	}
	queue := [][]string{{start}}
	for len(queue) > 0 {
		path := queue[0]
		queue = queue[1:]

		g.GetNeighbors(path, &queue, visited)
		fmt.Println(queue)
	}
}

func (g *GraphData) GetNeighbors(path []string, q *[][]string, v map[string]bool) {
	s := path[len(path)-1]
	for _, r := range g.Rooms {
		if r.Key == s {
			for _, nei := range r.Neighbors {
				if !v[nei.Key] {
					newpath := []string{}
					v[nei.Key] = true
					newpath = path
					newpath = append(newpath, nei.Key)
					*q = append(*q, newpath)
				}
			}
		}
	}
}
