package utils

import (
	"strconv"
	"strings"
)

func CheckDuplicates(lines []string) bool {
	for i := 0; i < len(lines); i++ {
		for j := i + 1; j < len(lines); j++ {
			if lines[i] == lines[j] {
				return false
			}
		}
	}

	return true
}

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
