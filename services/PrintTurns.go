package services

import (
	"fmt"
	"slices"
	"strings"
)

func allPathsEmpty(groups []AntsGroup) bool {
	for _, group := range groups {
		for _, path := range group.Paths {
			if len(path) > 0 {
				return false
			}
		}
	}
	return true
}

func PrintTurns(g []AntsGroup, er string) {
	Turns := [][]string{}
	turn := []string{}
	ar := []string{}
	awor := GetAntswithOneRoom(g)
	for !allPathsEmpty(g) {
		for j := 0; j < len(g); j++ {
			for i := 0; i < len(g[j].Ants); i++ {
				if len(g[j].Paths[i]) == 0 {
					continue
				}
				if slices.Contains(awor, g[j].Ants[i]) && NotinTurn(ar, awor) {
					continue
				}
				if !contains(turn, g[j].Paths[i][0]) {
					if g[j].Paths[i][0] != er {
						turn = append(turn, g[j].Paths[i][0])
					}
					s := fmt.Sprintf("L%d-%s", g[j].Ants[i], g[j].Paths[i][0])
					ar = append(ar, s)
					g[j].Paths[i] = g[j].Paths[i][1:]
				}
			}
		}
		if len(ar) > 0 {
			Turns = append(Turns, ar)
		}
		turn = []string{}
		ar = []string{}
	}

	for _, ele := range Turns {
		fmt.Println(strings.Join(ele, " "))
	}
}

func contains(slice []string, item string) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

func GetAntswithOneRoom(g []AntsGroup) []int {
	Ants := []int{}
	for _, ele := range g {
		for i, a := range ele.Paths {
			if len(a) == 1 {
				Ants = append(Ants, ele.Ants[i])
			}
		}
	}
	return Ants
}

func NotinTurn(ar []string, awor []int) bool {
	for i := 0; i < len(ar); i++ {
		for j := 0; j < len(awor); j++ {
			ants := fmt.Sprintf("L%d-", awor[j])
			if strings.HasPrefix(ar[i], ants) {
				return true
			}
		}
	}
	return false
}
