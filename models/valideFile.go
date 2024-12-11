package models

import (
	"bufio"
	"os"
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
	Path [][]string
}

// Create an instance from the struct GraphData
func NewGraphData() *GraphData {
	return &GraphData{
		Rooms:    make(map[string][]string),
		Tunneles: make(map[string][]string),
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
			g.NbOfants, err = strconv.Atoi(line)
			if err != nil {
				return "invalid number of ants"
			}

			if g.NbOfants <= 0 {
				return "invalid number of ants"
			}
			count = 1
		}

		if line == "" || (line[0] == '#' && line != "##start" && line != "##end") {
			continue
		}

		room := strings.Fields(line)
		if count == 2 || count == 3 {
			if !utils.IsValidRomm(line) {
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
				return "Invalid End"
			}
			count = 3
			continue
		}

		if utils.IsValidRomm(line) {
			g.Rooms[room[0]] = room[1:]
			continue
		}

		if utils.IsValidTunnel(line) {
			tunnel := strings.Split(line, "-")
			g.Tunneles[tunnel[0]] = append(g.Tunneles[tunnel[0]], tunnel[1])
			g.Tunneles[tunnel[1]] = append(g.Tunneles[tunnel[1]], tunnel[0])
		}

	}

	if g.Start == "" || g.End == "" || g.Start == g.End {
		return "Error start or end"
	}

	g.CheckLinks()

	return ""
}
