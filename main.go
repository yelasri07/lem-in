package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"lemin/functions"
)

func main() {
	file, err := os.Open("sample.txt")
	if err != nil {
		fmt.Println("File error")
		return
	}

	lines := []string{}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			lines = append(lines, line)
		}
	}

	if len(lines) == 0 {
		fmt.Println("error len equal 0")
		return
	}

	infos := functions.NewInfos()

	msg := infos.Treatment(lines)
	if msg != "" {
		fmt.Println(msg)
		return
	}

	fmt.Println(infos.NbOfants)
	fmt.Println(infos.Start)
	fmt.Println(infos.Rooms)
}
