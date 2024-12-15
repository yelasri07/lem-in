package services

import "slices"

func (g *GraphData) BFS(r string, v *map[string]bool) {
	g.BFSHelper(r, v)
}

func (g *GraphData) BFSHelper(room string, visited *map[string]bool) {
	var queue [][]string
	var currentPath []string
	(*visited)[g.Start] = true
	(*visited)[room] = true

	queue = append(queue, []string{g.Start, room})
	for len(queue) > 0 {

		currentPath = queue[0]

		queue = queue[1:]

		if string(currentPath[len(currentPath)-1]) == g.End {
			g.Paths = append(g.Paths, currentPath)
			UnlockRooms(currentPath, queue, *visited)
			break
		}

		for _, neighbor := range g.Tunnels[currentPath[len(currentPath)-1]] {
			if !(*visited)[neighbor] {
				newPath := append([]string{}, currentPath...)
				newPath = append(newPath, neighbor)
				queue = append(queue, newPath)
				if neighbor != g.End {
					(*visited)[neighbor] = true
				}
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
