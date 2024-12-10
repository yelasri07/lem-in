package main

import (
	"fmt"
	"os"

	"lemin/models"
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

	infos := models.NewGraphData()

	msg := infos.ValidateFileContent(file)
	if msg != "" {
		fmt.Println(msg)
		return
	}

	fmt.Println("Number of ants ==>", infos.NbOfants)
	fmt.Println("Start ==>", infos.Start)
	fmt.Println("End ==>", infos.End)
	fmt.Println("Rooms", infos.Rooms)
	fmt.Println("Tunnels==>", infos.Tunneles)
}
