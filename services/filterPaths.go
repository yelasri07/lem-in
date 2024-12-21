package services

func (g *GraphData) FilterPaths() {
	max := g.Groups[0]
	for _, grp := range g.Groups {
		if grp.lenPaths > max.lenPaths {
			max = grp
		}
	}

	g.Paths = append(g.Paths, max.key)
	g.Paths = append(g.Paths, max.Comb...)

	PrintSteps(g.Paths, g.NbOfAnts)
}
