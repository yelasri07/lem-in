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
	
	for _, p := range info.UniquePaths {
		
		info.FindMorePaths(p)
	}
	for _, p := range info.UniquePaths {
		fmt.Println("--------")
		fmt.Println(p)
		fmt.Println("---------")		
		info.FindGroups(p)
	}



	info.FindTheBestGroup()
	
}


