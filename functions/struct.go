package functions


type Info struct {
	NumberOfAnts int
	Start string
	End string
	Rooms map[string][]string
	Tunnels map[string][]string

	UniquePaths [][]string

	Neiofstart map[string]bool
	NumberOfrooms int
 	Res []string
} 	