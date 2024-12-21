package functions


type Info struct {
	NumberOfAnts int
	Start string
	End string
	
	Rooms map[string][]string
	Tunnels map[string][]string

	UniquePaths [][]string


 	Res []string
	NumberOfGroups int
	AllGroups [][][]string
	UniqueGroups [][][]string

}
