package services

import "fmt"

func PrintSteps(paths []*PathInfos, nbAnts int) {
	Ants := make([][]int, len(paths)) // had ants hiya nmel li ghadi ydouz me3a dak lpath wach fhemti wela mafhmtich

	for _, p := range paths {
		fmt.Println(p)
	}

	fmt.Println(Ants)

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
		Ants[k] = append(Ants[k], i)
	}

	fmt.Println(Ants)
}