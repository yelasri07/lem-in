package services

import (
	"fmt"
	"slices"
	"strings"
)

func PrintSteps(paths []*PathInfos, nbAnts int, end string) {

	ants := make([][]string, nbAnts)

	for _, p := range paths {
		fmt.Println(p)
	}

	for i := 1; i <= nbAnts; i++ {
		shortPath := paths[0]
		for j := 0; j < len(paths); j++ {
			if paths[j].len < shortPath.len {
				shortPath = paths[j]
			}
		}
		shortPath.len++

		ants[i-1] = append(ants[i-1], shortPath.Path...)
	}

	turns := [][]string{}
	turn := []string{}
	step := []string{}
	for !allPathsEmpty(ants) {

		for i := 0; i < len(ants); i++ {
			if len(ants[i]) == 0 {
				continue
			}

			if !slices.Contains(turn, ants[i][0]) {
				s := fmt.Sprintf("L%v-%v", i+1, ants[i][0])
				if ants[i][0] != end {
					turn = append(turn, ants[i][0])
				}
				ants[i] = ants[i][1:]
				step = append(step, s)
			}

		}

		turns = append(turns, step)

		turn = []string{}
		step = []string{}

	}

	for _, turn := range turns {
		fmt.Println(strings.Join(turn," "))
	}
}

func allPathsEmpty(groups [][]string) bool {
	for _, group := range groups {
		if len(group) > 0 {
			return false
		}
	}
	return true
}
