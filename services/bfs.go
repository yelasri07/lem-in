package services

import (
	"fmt"
	"slices"
)

func (g *GraphData) BFS() {
	g.BFSHelper()
}

func (g *GraphData) BFSHelper() {
	var queue [][]string
	var currentPath []string

	queue = append(queue, []string{g.Start})

	for len(queue) > 0 {

		currentPath = queue[0]

		queue = queue[1:]

		if string(currentPath[len(currentPath)-1]) == g.End {
			g.Paths = append(g.Paths, currentPath)
			continue
		}

		for _, neighbor := range g.Tunnels[currentPath[len(currentPath)-1]] {
			if !slices.Contains(currentPath, neighbor) {
				newPath := append([]string{}, currentPath...)
				newPath = append(newPath, neighbor)
				queue = append(queue, newPath)
			}
		}

	}
	fmt.Println(g.Paths)
}