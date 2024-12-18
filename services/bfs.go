package services

import (
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
		a := &Groups{key: path[:endRoom]}
		g.Groups = append(g.Groups, a)
	}
	for _, grp := range g.Groups {
		g.CombBfs(grp)
	}
}

func (g *GraphData) CombBfs(grp *Groups) {
	var queue [][]string
	var currentPath []string
	queue = append(queue, []string{g.Start})
	for len(queue) > 0 {
		currentPath = queue[0]
		queue = queue[1:]
		
		lastRoom := currentPath[len(currentPath)-1]
		if g.End == lastRoom {
			if Unique(grp, currentPath[1:]) {
				grp.Comb = append(grp.Comb, currentPath)
				continue
			}
		}

		for _, neighbor := range g.Tunnels[currentPath[len(currentPath)-1]] {
			if !slices.Contains(currentPath, neighbor) {
				newPath := append([]string{}, currentPath...)
				newPath = append(newPath, neighbor)
				queue = append(queue, newPath)

			}
		}

	}
}

func Unique(p *Groups, cp []string) bool {
	for i := 0; i < len(cp)-1; i++ {
		if slices.Contains(p.key, cp[i]) {
			return false
		}
		for j := 0; j < len(p.Comb); j++ {
			if slices.Contains(p.Comb[j], cp[i]) {
				return false
			}
		}
	}
	return true
}
