package main

import (
	"strings"
	"testing"
)

func TestBlockCreation(t *testing.T) {
	testLines := []string{
		"server { # php/fastcgi",
		"listen       80;",
		"server_name  domain1.com www.domain1.com;",
		"access_log   logs/domain1.access.log  main;",
		"root         html;",

		"location ~ \\.php$ {",
		"fastcgi_pass   127.0.0.1:1025;",
		"}",
		"}",
	}

	lines := getPointerArrayString(testLines)

	block := getAllBlocks(lines, 0, len(testLines), 0, -1)
	blockList := []*ngninxBlock{block}

	if !strings.Contains(*lines[block.Start-1], "{") {
		t.Errorf(*lines[block.Start-1]+" || %d", block.Start)
	}
	if !strings.Contains(*lines[block.End-1], "}") {
		t.Errorf(*lines[block.Start-1]+" || %d", block.Start)
	}

	if block.totalBlocks != 1 {
		t.Errorf("Nested Block Number Mismatch || %d", block.Start)
	}

	if !strings.Contains(*lines[block.NestedBlocks[0].End-1], "}") {
		t.Errorf(*lines[block.NestedBlocks[0].End-1]+" || %d", block.Start)
	}

	if !strings.Contains(*lines[block.NestedBlocks[0].Start-1], "{") {
		t.Errorf(*lines[block.NestedBlocks[0].Start-1]+" || %d", block.Start)
	}

	blocks := getNgnixBlocks(blockList, "www.domain1.com")
	var allLines *[]*string = blocks.allLines

	combString := ""
	for _, line := range *allLines {
		combString += *line + "\n"
	}

	if !strings.Contains(combString, "www.domain1.com") {
		t.Errorf(combString)
	}
}

func getPointerArrayString(array []string) []*string {
	var lines []*string
	for index, _ := range array {
		lines = append(lines, &array[index])
	}

	return lines
}
