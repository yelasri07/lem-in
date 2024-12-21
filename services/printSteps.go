package services

import (
	"fmt"
	"strings"
)

func PrintSteps(paths []*PathInfos, nbAnts int, er string) {
	Ants := make([]*PathInfos, nbAnts)
	var Grps []AntsGroup

	for _, p := range paths {
		fmt.Println(p.Path)
	}

	fmt.Println("-------------------")

	grpAnt := 0
	for i := 1; i <= nbAnts; i++ {
		petitPath := paths[0]
		k := 0
		for j := 0; j < len(paths); j++ {
			if paths[j].len < petitPath.len {
				petitPath = paths[j]
				k = j
			}
		}
		petitPath.len++
		Ants[i-1] = paths[k]
	}

	for i, c := range Ants {
		fmt.Printf("nemla %v flpath : %v\n", i+1, c.Path)
	}

	fmt.Println("-------------")

	a := []int{}
	p := [][]string{}
	for i, c := range Ants {
		if grpAnt != len(paths) {
			a = append(a, i+1)
			p = append(p, c.Path)
			grpAnt++
		}
		if grpAnt == len(paths) || i == len(Ants)-1 {
			st := &AntsGroup{Ants: a, Paths: p}
			Grps = append(Grps, *st)
			a = []int{}
			p = [][]string{}
			grpAnt = 0
		}
	}
	PrintTurns(Grps, er)
}

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

// Imprimer les tours
func PrintTurns(g []AntsGroup, er string) {
	Turns := [][]string{}
	turn := []string{}
	ar := []string{}

	// Boucle jusqu'à ce que tous les chemins soient vides
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
