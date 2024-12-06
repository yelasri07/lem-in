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

func NewInfos() *GraphData {
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
		if lines[i] == "##start" {

			continue
		}
		room := strings.Fields(lines[i])

		if utils.IsValidRomm(room) && g.Start == "" {
			g.Start = room[0]
			g.Rooms[room[0]] = room[1:]

		}

		if utils.IsValidRomm(room) {
			g.Rooms[room[0]] = room[1:]
		}
	}

	return ""
}
