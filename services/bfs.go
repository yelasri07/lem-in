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
			if g.intersectionPoint(currentPath) {
				path := Paths{len: len(currentPath), rooms: currentPath}
				g.Paths = append(g.Paths, path)
			}
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

	g.PrintSteps()
}

func (g *GraphData) FindBestPaths() {
	for _, path := range g.Paths {
		fmt.Println(strings.Join(path.rooms, " -> "), "|| len :", path.len)
	}

	// fmt.Println(g.Paths)
}

func (g *GraphData) intersectionPoint(currentPath []string) bool {
	for i := 1; i < len(currentPath)-1; i++ {
		for j := 0; j < len(g.Paths); j++ {
			if slices.Contains(g.Paths[j].rooms, currentPath[i]) {
				return false
			}
		}
	}

	return true
}
