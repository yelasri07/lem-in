package services

import (
	"fmt"
	"slices"
	"strings"
)

func PrintSteps(paths []*PathInfos, nbAnts int, er string) {
	Ants := make([]*PathInfos, nbAnts) // had ants hiya nmel li ghadi ydouz me3a dak lpath wach fhemti wela mafhmtich
	var Grps []AntsGroup

	grpAnt := 0
	// Turns := [][]string{}
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

func PrintTurns(g []AntsGroup, er string) {
	Turns := [][]string{}
	turn := []string{}
	ar := []string{}

	for len(g[len(g)-1].Paths[len(g[len(g)-1].Paths)-1]) > 0 {
		for j := 0; j < len(g); j++ {
			for i := 0; i < len(g[j].Ants); i++ {
				if len(g[j].Paths[i]) == 0 {
					continue
				}

				if !slices.Contains(turn, g[j].Paths[i][0]) {
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
