package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type block struct {
	Start        int
	End          int
	AllLines     *[]*string
	NestedBlocks []*block
	totalBlocks  int
}

func main() {
	file, _ := os.Open("ngnix.conf")
	scanner := bufio.NewScanner(file)
	var lines []string
	start := false

	for scanner.Scan() {
		currentLine := scanner.Text()
		if strings.Contains(currentLine, "{") || start {
			start = true
			lines = append(lines, currentLine)
		}
		// fmt.Println(scanner.Text())
	}

	fmt.Print(lines)
}
