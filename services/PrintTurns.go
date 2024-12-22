package services

import (
	"fmt"
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

	for !allPathsEmpty(g) {
		for j := 0; j < len(g); j++ {
			for i := 0; i < len(g[j].Ants); i++ {
				if len(g[j].Paths[i]) == 0 {
					continue
				}

				// Vérifie si l'étape actuelle est déjà occupée
				if !contains(turn, g[j].Paths[i][0]) {
					if g[j].Paths[i][0] != er {
						turn = append(turn, g[j].Paths[i][0])
					}
					s := fmt.Sprintf("L%d-%s", g[j].Ants[i], g[j].Paths[i][0])
					ar = append(ar, s)
					// Passe à l'étape suivante
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
