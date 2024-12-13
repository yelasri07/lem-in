package main

import (
	"fmt"
	"os"

	"lemin/services"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Enter file name")
		return
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("File error")
		return
	}
	defer file.Close()

	infos := services.NewGraphData()

	msg := infos.ValidateFileContent(file)
	if msg != "" {
		fmt.Println(msg)
		return
	}

	// fmt.Println("Number of ants ==>", infos.NbOfAnts)
	// fmt.Println("Start ==>", infos.Start)
	// fmt.Println("End ==>", infos.End)
	// for _, room := range infos.Rooms {
	// 	fmt.Printf("room: %v has neighbors: ", room.Key)
	// 	for _, neighbor := range room.Neighbors {
	// 		fmt.Printf("%v ", neighbor.Key)
	// 	}
	// 	fmt.Println("")
	// }
}
