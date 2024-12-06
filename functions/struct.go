package functions

import (
	"lemin/utils"
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

func (inf *Infos) ValidateFileContent(lines []string) string {
	var err error
	inf.NbOfants, err = strconv.Atoi(lines[0])
	if err != nil {
		return "invalid number of ants"
	}

	if inf.NbOfants <= 0 {
		return "Number of ants is negative"
	}

	for i := 1; i < len(lines); i++ {
		if lines[i] == "##start" {

			continue
		}
		room := strings.Fields(lines[i])

		if utils.IsValidRomm(room) && inf.Start == "" {
			inf.Start = room[0]
			inf.Rooms[room[0]] = room[1] + " " + room[2]

		}

		if utils.IsValidRomm(room) {
			inf.Rooms[room[0]] = room[1] + " " + room[2]
		}
	}

	return ""
}
