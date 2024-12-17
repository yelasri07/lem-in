package functions


func (y * Info) Hh(path []string) {
	var NewMatrix [][]string
	
	NewMatrix = append(NewMatrix, path)
	for i:=0 ; i < len(y.TheBestpaths); i++ {
		if Unique(y.TheBestpaths[i] , path) {
			NewMatrix = append(NewMatrix, y.TheBestpaths[i])
		}
	}
	
	y.All = append(y.All, NewMatrix)	
}


func Unique(p1, p2 []string) bool {
	for i := 0; i < len(p1)-1; i++ {
		for j := 0; j < len(p2)-1; j++ {
			if p1[i] == p2[j] {
				return false
			}
		}
	}
	return true
}