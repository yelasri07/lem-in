package services

import "fmt"

func PrintSteps(paths []*PathInfos, nbAnts int, er string) {
	Ants := make([]*PathInfos, nbAnts)
	var Grps []AntsGroup
	a := []int{}
	p := [][]string{}
	grpAnt := 0

	for _, p := range paths {
		fmt.Println(p)
	}

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

		if grpAnt != len(paths) {

			if i <= nbAnts {
				a = append(a, i)
				p = append(p, petitPath.Path)
			}
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
	if len(a) != 0 && len(p) != 0 {
		st := &AntsGroup{Ants: a, Paths: p}
		Grps = append(Grps, *st)
	}

	PrintTurns(Grps, er)
}

func VerifyDupPath(paths [][]string, p []string, er string) bool {
	for _, path := range paths {
		for i := 0; i < len(path)-1; i++ {
			for j := 0; j < len(p)-1; j++ {
				if path[i] == p[j] {
					return true
				}
			}
		}
	}
	return false
}
