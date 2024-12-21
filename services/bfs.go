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

// GroupMaker creates groups for each path by initializing a group for every path in g.Paths
// and then combining them using the CombBfs function.
func (g *GraphData) GroupMaker() {
	for _, path := range g.Paths {
		group := &Groups{key: path, lenPaths: 1}
		g.Groups = append(g.Groups, group)
	}

	g.Paths = []*PathInfos{}

	for _, grp := range g.Groups {
		g.CombBFS(grp)
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
				p := &PathInfos{len: len(currentPath), Path: currentPath[1:]}
				grp.Comb = append(grp.Comb, p)
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

// Unique checks if the given path (cp) shares any common room with the group's key path  
func Unique(p *Groups, currentPath []string) bool {
	for i := 0; i < len(currentPath)-1; i++ {
		if slices.Contains(p.key.Path, currentPath[i]) {
			return false
		}
		for j := 0; j < len(p.Comb); j++ {
			if slices.Contains(p.Comb[j].Path, currentPath[i]) {
				return false
			}
		}
	}
	return true
}
