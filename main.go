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
	}

	// fmt.Println(info.NumberOfAnts)
	// fmt.Println(info.Start)
	// fmt.Println(info.End)
	// fmt.Println(info.Rooms)
	// fmt.Println(info.Tunnels)

	for _, p := range info.UniquePaths {
		fmt.Println(p)
	}
}
