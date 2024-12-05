package functions

import (
	"strconv"
	"strings"
)

type Infos struct {
	NbOfants int
	Start    string
	End      string
	Rooms    map[string]string
	Tunneles map[string]string
}

func NewInfos() *Infos {
	return &Infos{
		Rooms:    make(map[string]string),
		Tunneles: make(map[string]string),
	}
}

func (inf *Infos) Treatment(lines []string) string {
	nbAnts := strings.Fields(lines[0])

	if len(nbAnts) != 1 {
		return "Error number ants"
	}

	nAnts, err := strconv.Atoi(nbAnts[0])
	if err != nil {
		return "invalid number of ants"
	}

	if nAnts <= 0 {
		return "Number of ants is negative"
	}

	inf.NbOfants = nAnts
	start := false
	lines = lines[1:]
	for i := 0; i < len(lines); i++ {
		if lines[i] == "##start" && inf.Start != "" {
			return "Invalid form of file."
		}

		if lines[i] == "##start" {
			start = true
			continue
		}
		room := strings.Fields(lines[i])

		if start && isValidRomm(room) && inf.Start == "" {
			inf.Start = room[0]
			inf.Rooms[room[0]] = room[1] + " " + room[2]
			start = false
		}

		if isValidRomm(room) {
			inf.Rooms[room[0]] = room[1] + " " + room[2]
		}

	}

	return ""
}

func isValidRomm(room []string) bool {
	if len(room) != 3 {
		return false
	}
	if room[0][0] == 'L' || room[0][0] == 'l' || room[0][0] == '#' {
		return false
	}
	_, err := strconv.Atoi(room[1])
	_, err2 := strconv.Atoi(room[2])
	if err != nil || err2 != nil {
		return false
	}
	return true
}
