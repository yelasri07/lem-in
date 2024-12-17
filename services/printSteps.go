package services

import "fmt"

func (g *GraphData) PrintSteps() {
	ants := 0

	for ants != g.NbOfAnts {
		ants++

		for i := 0 ; i < len(g.Paths) ; i++ {
			for j := 1 ; j < len(g.Paths[i].rooms) ; j++ {
				fmt.Printf("L%v-%v\n", ants, g.Paths[i].rooms[j])
			}
			
		}
	}
}
