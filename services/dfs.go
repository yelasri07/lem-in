package services

import (
	"fmt"
	"strings"
	"sync"
)

func (g *GraphData) DFS() {
	var startRoom *Room
	for _, room := range g.Rooms {
		if room.Key == g.Start {
			startRoom = room
		}
	}

	for _, a := range startRoom.Neighbors {
		g.Neiofstart = append(g.Neiofstart, a)
	}

	var wg sync.WaitGroup

	for _, nei := range startRoom.Neighbors {
		wg.Add(1)
		go func() {
			defer wg.Done()
			visited := make(map[string]bool)
		var currentPath []string
		g.DFSHelper(nei, visited, currentPath)
		}()
		
	}

	wg.Wait()
	

	for _, path := range g.Paths {
		fmt.Println(strings.Join(path, " -> "))
	}
}

// Recursive DFS helper function
func (g *GraphData) DFSHelper(room *Room, visited map[string]bool, currentPath []string) {
	visited[room.Key] = true
	currentPath = append(currentPath, room.Key)

	if room.Key == g.End {
		newPath := make([]string, len(currentPath))
		copy(newPath, currentPath)
		g.Paths = append(g.Paths, newPath)
	}

	for _, neighbors := range room.Neighbors {
		if !visited[neighbors.Key] && neighbors.Key != g.Start && isayn(g.Neiofstart, neighbors) {
			g.DFSHelper(neighbors, visited, currentPath)
		}
	}

	visited[room.Key] = false
}

func isayn(a []*Room, b *Room) bool {
	for _, nei := range a {
		if b == nei {
			return false
		}
	}

	return true
}
