package services

import (
	"bufio"
	"errors"
	"os"
	"strconv"
	"strings"

	"lemin/utils"
)

type GraphData struct {
	NbOfAnts int
	Start    string
	End      string
	Rooms    map[string][]string // Contains the room name and its coordinates.
	Tunnels  map[string][]string // Contains the room and its connected neighbors.
	Paths    []*PathInfos        // The slice contains structs, each representing information about a path.
	Groups   []*Groups           // The slice contains structs, each representing a group of paths.
}

type Groups struct {
	key      *PathInfos
	Comb     []*PathInfos
	lenPaths int
}

type PathInfos struct {
	len  int
	Path []string
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

	var count int
	for scanner.Scan() {

		line := strings.TrimSpace(scanner.Text())
		if count == 0 {
			g.NbOfAnts, err = strconv.Atoi(line)
			if err != nil {
				return "ERROR: invalid data format"
			}

			if g.NbOfAnts <= 0 {
				return "ERROR: invalid data format"
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
				return "ERROR: invalid data format"
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
				return "ERROR: invalid data format"
			}
			count = 2
			continue
		}

		if line == "##end" {
			if g.End != "" {
				return "ERROR: invalid data format"
			}
			count = 3
			continue
		}
		if !utils.IsValidRoom(line) && !utils.IsValidTunnel(line) {
			return "ERROR: invalid data format"
		}
		if utils.IsValidRoom(line) {
			err := g.AddRoom(line)
			if err != nil {
				return "ERROR: invalid data format"
			}
			continue
		}
		if utils.IsValidTunnel(line) {
			err := g.AddNeighbor(line)
			if err != nil {
				return "ERROR: invalid data format"
			}
		}

	}

	if g.Start == "" || g.End == "" || g.Start == g.End {
		return "ERROR: invalid data format"
	}

	for i := 0; i < len(g.Tunnels[g.Start]); i++ {
		g.BFS(g.Tunnels[g.Start][i])
	}

	g.GroupMaker()

	g.FilterPaths()

	return ""
}

// AddRoom adds a new room to the collection.
func (g *GraphData) AddRoom(line string) error {
	room := strings.Fields(line)
	if _, exist := g.Rooms[room[0]]; exist {
		return errors.New("ERROR: invalid data format")
	}

	g.Rooms[room[0]] = append(g.Rooms[room[0]], room[1], room[2])
	return nil
}

// AddNeighbor adds a new tunnel to the collection.
func (g *GraphData) AddNeighbor(line string) error {
	if !utils.ContainsRoom(line, g.Rooms) {
		return errors.New("ERROR: invalid data format")
	}

	tunnel := strings.Split(line, "-")
	g.Tunnels[tunnel[0]] = append(g.Tunnels[tunnel[0]], tunnel[1])
	g.Tunnels[tunnel[1]] = append(g.Tunnels[tunnel[1]], tunnel[0])
	return nil
}
