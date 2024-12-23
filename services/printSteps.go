package services

func PrintSteps(paths []*PathInfos, nbAnts int, end string) {
	var Grps []AntsGroup
	ants := []int{}
	currentPath := [][]string{}
	grpAntCount := 0

	for i := 1; i <= nbAnts; i++ {
		shortPath := paths[0]
		for j := 0; j < len(paths); j++ {
			if paths[j].len < shortPath.len {
				shortPath = paths[j]
			}
		}
		shortPath.len++

		if grpAntCount != len(paths) {
			ants = append(ants, i)
			currentPath = append(currentPath, shortPath.Path)
			grpAntCount++
		}

		if grpAntCount == len(paths) {
			st := &AntsGroup{Ants: ants, Paths: currentPath}
			Grps = append(Grps, *st)
			ants = []int{}
			currentPath = [][]string{}
			grpAntCount = 0
		}

	}
	if len(ants) != 0 && len(currentPath) != 0 {
		st := &AntsGroup{Ants: ants, Paths: currentPath}
		Grps = append(Grps, *st)
	}

	PrintTurns(Grps, end)
}
