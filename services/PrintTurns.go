package services

import (
	"fmt"
	"slices"
	"strings"
)

func allPathsEmpty(groups [][]string) bool {
	for _, group := range groups {
		if len(group) > 0 {
			return false
		}
	}
	return true
}

func (g *GraphData) PrintTurns(ants [][]string) {
	turns := [][]string{}
	turn := []string{}
	step := []string{}
	for !allPathsEmpty(ants) {

		badTunnel := false

		for i := 0; i < len(ants); i++ {
			if len(ants[i]) == 0 {
				continue
			}

			s := fmt.Sprintf("L%v-%v", i+1, ants[i][0])
			if len(ants[i]) == 2 && ants[i][0] == g.End && ants[i][1] == g.End {
				if !badTunnel {
					ants[i] = ants[i][2:]
					step = append(step, s)
					badTunnel = true
				}
			} else if !slices.Contains(turn, ants[i][0]) {
				if ants[i][0] != g.End {
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
		fmt.Println(strings.Join(turn, " "))
	}
}
