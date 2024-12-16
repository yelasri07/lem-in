package services

import (
	"fmt"
	"slices"
	"strings"
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

		lastRoomCurrent := currentPath[len(currentPath)-1]
		if lastRoomCurrent == g.End {
			g.Paths = append(g.Paths, currentPath)
			continue
		}

		for _, neighbor := range g.Tunnels[lastRoomCurrent] {
			if !slices.Contains(currentPath, neighbor) {
				newPath := append([]string{}, currentPath...)
				newPath = append(newPath, neighbor)
				queue = append(queue, newPath)
			}
		}

	}

	g.FindBestPaths()

}

func (g *GraphData) FindBestPaths() {
	for _, path := range g.Paths {
		fmt.Println(strings.Join(path, " -> "))
	}
}