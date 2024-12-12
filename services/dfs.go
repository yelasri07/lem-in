package services

import (
	"fmt"
)

func (g *GraphData) DFS() {
	visited := make(map[string]bool) // Map to track visited rooms

	// Find the start room
	startRoom := g.getRoomByKey(g.Start)
	if startRoom == nil {
		fmt.Println("Start room not found.")
		return
	}

	// Perform DFS
	var currentPath []string
	g.dfsHelper(startRoom, visited, currentPath)

	// fmt.Println(g.Paths)
}

// Recursive DFS helper function
func (g *GraphData) dfsHelper(room *Room, visited map[string]bool, currentPath []string) {
	// Mark the current room as visited

	visited[room.Key] = true
	currentPath = append(currentPath, room.Key) // Add current room to path

	// Check if this is the end room

	if room.Key == g.End {
		fmt.Println(currentPath)
		g.Paths = append(g.Paths, currentPath)
		fmt.Println(g.Paths)

	}

	// Visit all unvisited neighborss
	for _, neighbor := range room.Neighbors {
		if !visited[neighbor.Key] {
			g.dfsHelper(neighbor, visited, currentPath)
		}
	}

	// Backtrack: unmark it as visited
	if room.Key == g.End {
		visited[room.Key] = true
	} else {
		visited[room.Key] = false
	}
}

// Helper function to find a room by its key
func (g *GraphData) getRoomByKey(key string) *Room {
	for _, room := range g.Rooms {
		if room.Key == key {
			return room
		}
	}
	return nil
}
