package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"lemin/models"
	"lemin/utils"
)

func main() {
	file, err := os.Open("sample.txt")
	if err != nil {
		fmt.Println("File error")
		return
	}
	defer file.Close()

	lines := []string{}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || (line[0] == '#' && line != "##start" && line != "##end") {
			continue
		}
		lines = append(lines, line)
	}

	if !utils.CheckDuplicates(lines) {
		fmt.Println("error data duplicates")
		return
	}

	if len(lines) < 6 {
		fmt.Println("error len lines")
		return
	}

	infos := models.NewGraphData()

	msg := infos.ValidateFileContent(lines)
	if msg != "" {
		fmt.Println(msg)
		return
	}

	fmt.Println("Number of ants ==>",infos.NbOfants)
	fmt.Println("Start ==>",infos.Start)
	fmt.Println("End ==>",infos.End)
	fmt.Println(infos.Rooms)
}
