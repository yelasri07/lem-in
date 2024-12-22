package services

func PrintSteps(paths []*PathInfos, nbAnts int, er string) {
	Ants := make([]*PathInfos, nbAnts)
	var Grps []AntsGroup
	a := []int{}
	p := [][]string{}
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
		if grpAnt != len(paths) {
			for _, ele := range paths {
				if VerifyDupPath(p, ele.Path, er) {
					continue
				}
				a = append(a, i)
				p = append(p, ele.Path)
				grpAnt++
				break
			}
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

func VerifyDupPath(paths [][]string, p []string, er string) bool {
	for _, path := range paths {
		for i := 0; i < len(path)-1; i++ { // Ignore le dernier élément de 'path'
			for j := 0; j < len(p)-1; j++ { // Ignore le dernier élément de 'p'
				if path[i] == p[j] { // Vérifie si les éléments sont identiques
					return true
				}
			}
		}
	}
	return false
}
