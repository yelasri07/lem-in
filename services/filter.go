package services

import "slices"

// GroupMaker creates groups for each path by initializing a group for every path in g.Paths
// and then combining them using the CombBfs function.
func (g *GraphData) GroupMaker() {
	for _, path := range g.Paths {
		group := &Groups{key: path, lenPaths: 1}
		g.Groups = append(g.Groups, group)
	}

	g.Paths = []*PathInfos{}

	for _, grp := range g.Groups {
		g.CombBFS(grp)
	}
}

// Unique checks if the given path (currentPath) shares any common room with the group's key path
func Unique(p *Groups, currentPath []string) bool {
	if len(currentPath) == 1 {
		return false
	}
	for i := 0; i < len(currentPath)-1; i++ {
		if slices.Contains(p.key.Path, currentPath[i]) {
			return false
		}
		for j := 0; j < len(p.Comb); j++ {
			if slices.Contains(p.Comb[j].Path, currentPath[i]) {
				return false
			}
		}
	}
	return true
}

// FilterPaths give the max len group
func (g *GraphData) FilterPaths() {
	max := g.Groups[0]
	for _, grp := range g.Groups {
		if grp.lenPaths > max.lenPaths {
			max = grp
		}
	}

	g.Paths = append(g.Paths, max.key)
	g.Paths = append(g.Paths, max.Comb...)
}
