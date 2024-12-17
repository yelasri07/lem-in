package services

import (
	"fmt"
	"slices"
)

func (g *GraphData) BFS(r string) {
	g.BFSHelper(r)
	
}

func (g *GraphData) BFSHelper(room string) {
	var queue [][]string
	var currentPath []string
	visited := make(map[string]bool)
	visited[g.Start] = true
	visited[room] = true

	queue = append(queue, []string{g.Start, room})
	for len(queue) > 0 {

		currentPath = queue[0]
		queue = queue[1:]

		if string(currentPath[len(currentPath)-1]) == g.End {
			g.Paths = append(g.Paths, currentPath)
			break
		}

		for _, neighbor := range g.Tunnels[currentPath[len(currentPath)-1]] {
			if !visited[neighbor] {
				newPath := append([]string{}, currentPath...)
				newPath = append(newPath, neighbor)
				queue = append(queue, newPath)
				if neighbor != g.End {
					visited[neighbor] = true
				}
			}
		}
	}
}

func (g *GraphData) SortPath() {
	fmt.Println(g.Paths)
}

func UnlockRooms(path []string, inVistedRoom [][]string, visited map[string]bool) {
	for _, p := range inVistedRoom {
		for _, r := range p[2:] {
			if !slices.Contains(path, r) {
				visited[r] = false
			}
		}
	}
}
