package models

import (
	"lemin/utils"
	"strconv"
	"strings"
)

type GraphData struct {
	NbOfants int
	Start    string
	End      string
	Rooms    map[string][]string
	Tunneles map[string]string
}

func NewGraphData() *GraphData {
	return &GraphData{
		Rooms:    make(map[string][]string),
		Tunneles: make(map[string]string),
	}
}

func (g *GraphData) ValidateFileContent(lines []string) string {
	var err error
	g.NbOfants, err = strconv.Atoi(lines[0])
	if err != nil {
		return "invalid number of ants"
	}

	if g.NbOfants <= 0 {
		return "Number of ants is negative"
	}

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

		room := strings.Fields(lines[i])

		if utils.IsValidRomm(lines[i]) {
			if _,ok := g.Rooms[room[0]]; ok{
				return "error duplicate room name"
			}
			g.Rooms[room[0]] = room[1:]
		}
	}

	return ""
}
