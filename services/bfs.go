package services

import (
	"slices"
	"strings"
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

	queue = append(queue, []string{room})
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
	for i := 0; i <= len(g.Paths)-1; i++ {
		for j := i + 1; j < len(g.Paths); j++ {
			if len(g.Paths[i]) > len(g.Paths[j]) {
				g.Paths[i], g.Paths[j] = g.Paths[j], g.Paths[i]
			}
		}
	}
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

func (g *GraphData) GroupMaker() {
	for _, path := range g.Paths {
		endRoom := len(path) - 1
		g.CombBfs(path[:endRoom])
	}
}

func (g *GraphData) CombBfs(p []string) {
	var queue [][]string
	var currentPath []string
	visited := make(map[string]bool)
	queue = append(queue, []string{g.Start})
	visited[g.Start] = true
	for len(queue) > 1 {
		currentPath = queue[0]
		queue = queue[1:]
		lastRoom := currentPath[len(currentPath)-1]
		if g.End == lastRoom {
			if Unique(p, currentPath[1:]) {
			}
		}
	}
}

func Unique(p, cp []string) bool {
	for i := 0; i < len(cp)-1; i++ {
		if slices.Contains(p, cp[i]) {
			return true
		}
	}
	return false
}
