package services

import (
	"fmt"
	"slices"
)

func PrintSteps(paths []*PathInfos, nbAnts int) {
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
	PrintTurns(Grps)
}

func PrintTurns(g []AntsGroup) {
	Turns := [][]string{}
	turn := []string{}

	for i, group := range g {
		for j, grp2 := range g {
			if i!=j{
				if !slices.Contains(turn,)
			}
		}
	}
	fmt.Println(turn)
	fmt.Println(Turns)
}
