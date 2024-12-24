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

	/////////////

	info.Bfs(info.Start)

	if len(info.UniquePaths) == 1 {
		info.UniquePaths[0] = info.UniquePaths[0][1:]
	}

	info.Res = append(info.Res, info.Tunnels[info.Start]...)

	for _, v := range info.Tunnels[info.Start] {
		info.Bfs(v)
	}

	for _, p := range info.UniquePaths {
		info.FindMorePaths(p)
	}

	
	var res []string

	for _, p := range info.UniquePaths {
		if !info.StartWithEnd && len(p) == 1 {
			info.StartWithEnd = true
			res = p
			info.AllGroups = append(info.AllGroups, [][]string{p})
		}
		info.FindGroups(p)
	}
	

	if info.StartWithEnd {

		var varrr [][][]string
		varrr = append(varrr, [][]string{res})
		varrr = append(varrr, info.AllGroups...)

		info.AllGroups = varrr
	}

	info.FindTheBestGroup()


}
