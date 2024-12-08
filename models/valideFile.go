package models

import (
	"strconv"
	"strings"

	"lemin/utils"
)

type GraphData struct {
	NbOfants int
	Start    string
	End      string
	Rooms    map[string][]string
	Tunneles map[string][]string
}

// Create an instance from the struct GraphData
func NewGraphData() *GraphData {
	return &GraphData{
		Rooms:    make(map[string][]string),
		Tunneles: make(map[string][]string),
	}
}

// ValidateFileContent is a method that validates the content inside the file
func (g *GraphData) ValidateFileContent(lines []string) string {
	var err error
	g.NbOfants, err = strconv.Atoi(lines[0])
	if err != nil {
		return "invalid number of ants"
	}

	if g.NbOfants <= 0 {
		return "Number of ants is negative"
	}
	allRoomsFinded := false
	for i := 1; i < len(lines); i++ {
		if lines[len(lines)-1] == "##start" || lines[len(lines)-1] == "##end" {
			return "there is no start or end"
		}

		if lines[i] == "##start" || lines[i] == "##end" {
			if !utils.IsValidRomm(lines[i+1]) {
				return "error room start or end"
			}
			room := strings.Fields(lines[i+1])[0]
			if lines[i] == "##start" {
				g.Start = room
			} else {
				g.End = room
			}
		}
		if allRoomsFinded && !utils.IsValidTunnel(lines[i], g.Rooms) {
			return "error format of file or tunnels data.."
		}
		if utils.IsValidTunnel(lines[i], g.Rooms) && !allRoomsFinded {
			allRoomsFinded = true
		}

		room := strings.Fields(lines[i])
		if utils.IsValidRomm(lines[i]) {
			if _, ok := g.Rooms[room[0]]; ok {
				return "error duplicate room name"
			}
			g.Rooms[room[0]] = room[1:]
		}
		if allRoomsFinded && utils.IsValidTunnel(lines[i], g.Rooms) {
			tunnel := strings.Split(lines[i], "-")
			g.Tunneles[tunnel[0]] = append(g.Tunneles[tunnel[0]], tunnel[1])
			g.Tunneles[tunnel[1]] = append(g.Tunneles[tunnel[1]], tunnel[0])
		}
	}

	return ""
}
