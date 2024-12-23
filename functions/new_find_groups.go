package functions

func (y *Info) FindGroups(path []string) {
	var NewMatrix [][]string

	NewMatrix = append(NewMatrix, path)

	for i := 0; i < len(y.UniquePaths); i++ {
		if unique(NewMatrix, y.UniquePaths[i]) {
			NewMatrix = append(NewMatrix, y.UniquePaths[i])
		}
	}

	y.AllGroups = append(y.AllGroups, NewMatrix)
}

func unique(mat [][]string, p []string) bool {
	mat = append(mat, p)

	mymap := make(map[string]bool)

	for _, p := range mat {
		for i := 0; i < len(p)-1; i++ {
			if !mymap[p[i]] {
				mymap[p[i]] = true
			} else {
				return false
			}
		}
	}
	return true
}
