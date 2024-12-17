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



	
	info.Jj()



    for i := 0; i < len(info.UniquePaths); i++ {
		info.Hh(info.UniquePaths[i])
	}
	


	for _, o := range info.All {
		fmt.Println(o)
		fmt.Println("<--------------------------->")
	}



	fmt.Println("<--------------------------->")

	for _, p := range info.UniquePaths {
		fmt.Println(p)
	}

	fmt.Println("<--------------------------->")

	for _, p := range info.TheBestpaths {
		fmt.Println(p)
	}
	
}


