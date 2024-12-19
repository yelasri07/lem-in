package services

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"lemin/utils"
)

type GraphData struct {
	NbOfAnts int
	Start    string
	End      string
	Rooms    map[string][]string
	Tunnels  map[string][]string
	Paths    [][]string
	Groups   []*Groups
}

type Groups struct {
	key  []string
	Comb [][]string
}

// Create an instance from the struct GraphData
func NewGraphData() *GraphData {
	return &GraphData{
		Rooms:   make(map[string][]string),
		Tunnels: make(map[string][]string),
	}
}

// ValidateFileContent is a method that validates the content inside the file
func (g *GraphData) ValidateFileContent(file *os.File) string {
	var err error

	scanner := bufio.NewScanner(file)
	lineError := 0
	var count int
	for scanner.Scan() {
		lineError++
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
			return "Invalid format of room, tunnel, or coordinates in line " + strconv.Itoa(lineError)
		}
		if utils.IsValidRoom(line) {
			g.AddRoom(line)
			continue
		}
		if utils.IsValidTunnel(line) {
			g.AddNeighbor(line)
		}

	}

	if g.Start == "" || g.End == "" || g.Start == g.End {
		return "Error start or end"
	}
	for i := 0; i < len(g.Tunnels[g.Start]); i++ {
		g.BFS(g.Tunnels[g.Start][i])
	}

	g.SortPath()

	for index, path := range g.Paths {
		if index != 0 {
			for i := 0; i <= len(g.Paths[0])-2; i++ {
				for j := 0; j < len(path)-1; j++ {
					if g.Paths[0][i] == path[j] {
						if len(path[:j]) > 2 {
							fmt.Println(g.Paths[0][i], "=>", j)
						}
					}
				}
			}
		}
	}

	g.GroupMaker()

	// g.Sendants(g.Paths)
	for _, grp := range g.Groups {
		fmt.Println(grp)
	}
	return ""
}

func (g *GraphData) AddRoom(line string) {
	room := strings.Fields(line)
	if _, exist := g.Rooms[room[0]]; exist {
		log.Fatal("Room already exists, you cannot duplicate rooms")
	}
	g.Rooms[room[0]] = append(g.Rooms[room[0]], room[1], room[2])
}

func (g *GraphData) AddNeighbor(line string) {
	if !utils.ContainsRoom(line, g.Rooms) {
		log.Fatal("You might be trying to add a non-existent room")
	}

	tunnel := strings.Split(line, "-")
	g.Tunnels[tunnel[0]] = append(g.Tunnels[tunnel[0]], tunnel[1])
	g.Tunnels[tunnel[1]] = append(g.Tunnels[tunnel[1]], tunnel[0])
}
