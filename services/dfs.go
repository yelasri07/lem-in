package services

// import (
// 	"fmt"
// 	"strings"
// 	"sync"
// )

// func (g *GraphData) DFS() {
// 	var startRoom *Room
// 	for _, room := range g.Rooms {
// 		if room.Key == g.Start {
// 			startRoom = room
// 		}
// 	}

// 	g.Neiofstart = startRoom.Neighbors

// 	var wg sync.WaitGroup

// 	for _, nei := range startRoom.Neighbors {
// 		wg.Add(1)
// 		go func() {
// 			defer wg.Done()
// 			visited := make(map[string]bool)
// 			var currentPath []string
// 			g.DFSHelper(nei, visited, currentPath)
// 		}()

// 	}

// 	wg.Wait()

// 	for _, path := range g.Paths {
// 		fmt.Println("0 ->", strings.Join(path, " -> "))
// 	}
// }

// // Recursive DFS helper function
// func (g *GraphData) DFSHelper(room *Room, visited map[string]bool, currentPath []string) {
// 	visited[room.Key] = true
// 	currentPath = append(currentPath, room.Key)

// 	if room.Key == g.End {
// 		newPath := make([]string, len(currentPath))
// 		copy(newPath, currentPath)
// 		g.Paths = append(g.Paths, newPath)
// 	}

// 	for _, neighbors := range room.Neighbors {
// 		if !visited[neighbors.Key] && neighbors.Key != g.Start && !isNeighborStart(g.Neiofstart, neighbors) {
// 			g.DFSHelper(neighbors, visited, currentPath)
// 		}
// 	}

// 	visited[room.Key] = false
// }

// func isNeighborStart(a []*Room, b *Room) bool {
// 	for _, nei := range a {
// 		if b == nei {
// 			return true
// 		}
// 	}

// 	return false
// }
