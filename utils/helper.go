package utils

import (
	"strconv"
	"strings"
)

// Check for duplicates inside the file

// Check if the room is valid or not
func IsValidRomm(str string) bool {
	room := strings.Fields(str)
	if len(room) != 3 {
		return false
	}
	if room[0][0] == 'L' || room[0][0] == 'l' {
		return false
	}
	_, err := strconv.Atoi(room[1])
	_, err2 := strconv.Atoi(room[2])
	if err != nil || err2 != nil {
		return false
	}
	return true
}

func IsValidTunnel(str string) bool {
	tunnel := strings.Split(str, "-")
	if len(tunnel) != 2 {
		return false
	}

	return true
}

func ContainsRoom(str string, rooms map[string][]string) bool {
	tunnel := strings.Split(str, "-")
	for _, room := range tunnel {
		if _, exist := rooms[room]; !exist {
			return false
		}
	}
	return true
}
