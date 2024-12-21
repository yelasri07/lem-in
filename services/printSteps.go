package services

import "fmt"

func PrintSteps(paths []*PathInfos, nbAnts int) {
	Ants := make([]*PathInfos, nbAnts)

	for _, p := range paths {
		fmt.Println(p.Path)
	}

	fmt.Println("-------------------")

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
		Ants[i-1] = paths[k]
	}

	for i, c := range Ants {
		fmt.Printf("nemla %v flpath : %v\n", i+1, c.Path)
	}

	fmt.Println("-------------")
}
