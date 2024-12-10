package models

import "log"

func (g *GraphData) CheckLinks() {
	contain := true

	for key := range g.Tunneles {
		if _, exist := g.Rooms[key]; !exist {
			contain = false
		}
	}

	if !contain {
		log.Fatal("The room given in the tunnels doesn't appear on the rooms map.")
	}
}
