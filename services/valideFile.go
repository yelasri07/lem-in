package services

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"

	"lemin/utils"
)

type GraphData struct {
	NbOfAnts   int
	Start      string
	End        string
	Rooms      []*Room
	Paths      [][]string
	Neiofstart []*Room
}

type Room struct {
	Key       string
	X         string
	Y         string
	Neighbors []*Room
}

// Create an instance from the struct GraphData
func NewGraphData() *GraphData {
	return &GraphData{}
}

// ValidateFileContent is a method that validates the content inside the file
func (g *GraphData) ValidateFileContent(file *os.File) string {
	var err error

	scanner := bufio.NewScanner(file)
	ln := 0
	var count int
	for scanner.Scan() {
		ln++
		line := strings.TrimSpace(scanner.Text())
		if count == 0 {
			g.NbOfAnts, err = strconv.Atoi(line)
			if err != nil {
				return "invalid number of ants"
			}

			if g.NbOfAnts <= 0 {
				return "invalid number of ants"
			}
			count = 1
			continue
		}

		if line == "" || (line[0] == '#' && line != "##start" && line != "##end") {
			continue
		}

		room := strings.Fields(line)
		if count == 2 || count == 3 {
			if !utils.IsValidRoom(line) {
				return "Invalid start or end"
			}
			if count == 2 {
				g.Start = room[0]
			} else {
				g.End = room[0]
			}
			count = 1
		}

		if line == "##start" {
			if g.Start != "" {
				return "Invalid start"
			}
			count = 2
			continue
		}

		if line == "##end" {
			if g.End != "" {
				return "Invalid end"
			}
			count = 3
			continue
		}
		if !utils.IsValidRoom(line) && !utils.IsValidTunnel(line) {
			return "Invalid format of room, tunnel, or coordinates in line " + strconv.Itoa(ln)
		}
		if utils.IsValidRoom(line) {
			g.AddRoom(line)
			continue
		}
		if utils.IsValidTunnel(line) {
			tunnel := strings.Split(line, "-")
			g.AddNeighbor(tunnel[0], tunnel[1])
		}

	}

	if g.Start == "" || g.End == "" || g.Start == g.End {
		return "Error start or end"
	}

	// g.DFS()
	g.Bfs()
	return ""
}

func (g *GraphData) AddRoom(line string) {
	room := strings.Fields(line)
	if contains(g.Rooms, room[0]) {
		log.Fatal("Room already exists, you cannot duplicate rooms")
	}
	g.Rooms = append(g.Rooms, &Room{Key: room[0], X: room[1], Y: room[2]})
}

func contains(rooms []*Room, key string) bool {
	for _, r := range rooms {
		if r.Key == key {
			return true
		}
	}
	return false
}

func (g *GraphData) AddNeighbor(from string, to string) {
	fromVertex := g.GetRoom(from)
	toVertex := g.GetRoom(to)
	if fromVertex == nil || toVertex == nil {
		log.Fatal("You might be trying to add a non-existent room")
	}
	if contains(fromVertex.Neighbors, to) {
		log.Fatal("You cannot add a neighbor that already exists")
	}
	fromVertex.Neighbors = append(fromVertex.Neighbors, toVertex)
	toVertex.Neighbors = append(toVertex.Neighbors, fromVertex)
}

func (g *GraphData) GetRoom(key string) *Room {
	for _, room := range g.Rooms {
		if room.Key == key {
			return room
		}
	}
	return nil
}
