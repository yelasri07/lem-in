package services

import "fmt"

func PrintSteps(paths []*PathInfos, nbAnts int, er string) {
	var Grps []AntsGroup
	ant := []int{}
	p := [][]string{}
	grpAnt := 0

	for i := 1; i <= nbAnts; i++ {
		petitPath := paths[0]
		for j := 0; j < len(paths); j++ {
			if paths[j].len < petitPath.len {
				petitPath = paths[j]
			}
		}
		petitPath.len++

		if grpAnt != len(paths) {
			if i <= nbAnts {
				ant = append(ant, i)
				p = append(p, petitPath.Path)
			}
			grpAnt++
		}

		if grpAnt == len(paths) {
			st := &AntsGroup{Ants: ant, Paths: p}
			Grps = append(Grps, *st)
			ant = []int{}
			p = [][]string{}
			grpAnt = 0
		}

	}
	if len(ant) != 0 && len(p) != 0 {
		st := &AntsGroup{Ants: ant, Paths: p}
		Grps = append(Grps, *st)
	}
	fmt.Println(Grps)
	PrintTurns(Grps, er)
}
