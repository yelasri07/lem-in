package services

func (g *GraphData) assignAntPaths() {
	ants := make([][]string, g.NbOfAnts)

	for i := 1; i <= g.NbOfAnts; i++ {
		shortPath := g.Paths[0]
		for j := 0; j < len(g.Paths); j++ {
			if g.Paths[j].len < shortPath.len {
				shortPath = g.Paths[j]
			}
		}
		shortPath.len++

		ants[i-1] = append(ants[i-1], shortPath.Path...)
	}

	for i := 0; i < len(ants); i++ {
		if len(ants[i]) == 1 {
			ants[i] = append(ants[i], g.End)
		}
	}

	g.PrintTurns(ants)
}
