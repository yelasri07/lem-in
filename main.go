package main

import (
	"fmt"
	"os"

	"lemin/functions"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	info := functions.Info{}

	err := info.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	
	
	
	info.Res = append(info.Res, info.Tunnels[info.Start]...)
	
	
	for _, v := range info.Tunnels[info.Start] {
		info.Bfs(v)
		if len(info.UniquePaths) >= 2 {
			info.FindGroups()
		}
	}
	
	//info.Bfs(info.Start)

	// fmt.Println(info.NumberOfAnts)
	// fmt.Println(info.Start)
	// fmt.Println(info.End)
	// fmt.Println(info.Rooms)
	//fmt.Println(info.Tunnels)

	for _, p := range info.UniquePaths {
		fmt.Println(p)
	}
	fmt.Println("<--------------------------->")
	for _, p := range info.TheBestpaths {
		fmt.Println(p)
	}
}
