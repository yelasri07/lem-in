package services

import (
	"fmt"
	"slices"
	"strings"
)

func (g *GraphData) BFS() {
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

		for _, room := range g.Rooms {
			if room.Key == string(currentPath[len(currentPath)-1]) {
				for _, neighbor := range room.Neighbors {
					if !slices.Contains(currentPath, neighbor.Key) {
						newPath := append([]string{},currentPath...)
						newPath = append(newPath, neighbor.Key)
						queue = append(queue, newPath)
					}
				
				}
				
				break
			}
			
		}

		fmt.Println(queue)

		// queue = [][]string{}

	}

	for _, path := range g.Paths {
		fmt.Println(strings.Join(path, " -> "))
	}
}
