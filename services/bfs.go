package services

import (
	"slices"
)

// bfs function implements the Breadth-First Search algorithm to find paths.
func (g *GraphData) BFS(neighborStart string) {
	var queue [][]string
	var currentPath []string

	queue = append(queue, []string{g.Start, neighborStart})
	for len(queue) > 0 {
		currentPath = queue[0]
		queue = queue[1:]
		lastRoom := currentPath[len(currentPath)-1]
		if lastRoom == g.End {
			path := &PathInfos{len: len(currentPath), Path: currentPath[1:]}
			g.Paths = append(g.Paths, path)
			break
		}

		for _, neighbor := range g.Tunnels[lastRoom] {
			if !slices.Contains(currentPath, neighbor) {
				newPath := append([]string{}, currentPath...)
				newPath = append(newPath, neighbor)
				queue = append(queue, newPath)
			}
		}
	}
}

// CombBfs finds additional paths and adds them to the group if they match its criteria.
func (g *GraphData) CombBFS(grp *Groups) {
	var queue [][]string
	var currentPath []string
	queue = append(queue, []string{g.Start})

	for len(queue) > 0 {
		currentPath = queue[0]
		queue = queue[1:]
		lastRoom := currentPath[len(currentPath)-1]
		if g.End == lastRoom {
			if Unique(grp, currentPath[1:]) {
				path := &PathInfos{len: len(currentPath), Path: currentPath[1:]}
				grp.Comb = append(grp.Comb, path)
				grp.lenPaths++
				continue
			}
		}

		for _, neighbor := range g.Tunnels[lastRoom] {
			if !slices.Contains(currentPath, neighbor) {
				newPath := append([]string{}, currentPath...)
				newPath = append(newPath, neighbor)
				queue = append(queue, newPath)

			}
		}

	}
}
