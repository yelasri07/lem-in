package functions

type Info struct {
	NumberOfAnts int
	Start        string
	End          string

	Rooms   map[string][]string
	Tunnels map[string][]string

	UniquePaths [][]string

	Res []string
	StartWithEnd bool
	NumberOfGroups int
	AllGroups    [][][]string
	UniqueGroups [][][]string
	SG int
	GG int
}

