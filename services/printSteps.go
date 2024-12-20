package services

import (
	"fmt"
	"strconv"
)

func PrintSteps(paths []*PathInfos, nbAnts int) {
	Ants := make([][]int, len(paths)) // had ants hiya nmel li ghadi ydouz me3a dak lpath wach fhemti wela mafhmtich

	for _, p := range paths {
		fmt.Println(p)
	}

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

	printAnts := make(map[string]*PathInfos)

	for i := 0; i < len(Ants); i++ {
		for j := 0; j < len(Ants[i]); j++ {
			printAnts["L"+strconv.Itoa(Ants[i][j])] = paths[i]
		}
	}

	fmt.Println(Ants)

	for a, p := range printAnts {
		fmt.Printf("nemla %v lpath li ghadi tmchi fih : %v\n", a , p.Path)
	}
}
