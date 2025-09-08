// Package pretty provides utilities for beautifying console output.
package pretty

import (
	"fmt"
	"log"
	"os"
	"strings"
)

// PrintProgress draw a progress bar in the terminal. The percent takes values from 0 to 1.
func PrintProgress(percent float64) {
	if percent < 0 || percent > 1 {
		log.Panicln("pretty: the percent takes values from 0 to 1")
	}
	out, err := os.Stdout.Stat()
	if err != nil {
		log.Panicln("pretty: cannot stat stdout:", err)
	}
	// Identify if we are displaying to a terminal or through a pipe or redirect.
	if out.Mode()&os.ModeCharDevice == os.ModeCharDevice {
		// Save or restore cursor position.
		if percent == 0 {
			log.Writer().Write([]byte{0x1b, 0x37})
		}
		if percent != 0 {
			log.Writer().Write([]byte{0x1b, 0x38})
		}
	}
	cap := int(percent * 44)
	buf := []byte("[                                             ] 000%")
	for i := 1; i < cap+1; i++ {
		buf[i] = '='
	}
	buf[1+cap] = '>'
	num := fmt.Sprintf("%3d", int(percent*100))
	buf[48] = num[0]
	buf[49] = num[1]
	buf[50] = num[2]
	log.Println("pretty:", string(buf))
}

// PrintTable easily draw tables in terminal/console applications from a list of lists of strings.
func PrintTable(data [][]string) {
	size := make([]int, len(data[0]))
	for _, r := range data {
		for j, c := range r {
			size[j] = max(size[j], len(c))
		}
	}
	line := make([]string, len(data[0]))
	for j, c := range data[0] {
		l := size[j]
		line[j] = c + strings.Repeat(" ", l-len(c))
	}
	log.Println("pretty:", strings.Join(line, " "))
	for i, c := range size {
		line[i] = strings.Repeat("-", c)
	}
	log.Println("pretty:", strings.Join(line, "-"))
	for _, r := range data[1:] {
		for j, c := range r {
			l := size[j]
			line[j] = c + strings.Repeat(" ", l-len(c))
		}
		log.Println("pretty:", strings.Join(line, " "))
	}
}

// Tree represents a node in a tree structure.
type Tree struct {
	Value string
	Elems []*Tree
}

func printTree(tree *Tree, prefix string) {
	for i, elem := range tree.Elems {
		isLast := i == len(tree.Elems)-1
		branch := "├── "
		if isLast {
			branch = "└── "
		}
		log.Println("pretty:", prefix+branch+elem.Value)
		if len(elem.Elems) > 0 {
			middle := "│   "
			if isLast {
				middle = "    "
			}
			printTree(elem, prefix+middle)
		}
	}
}

// PrintTree prints the tree structure starting from the root node.
func PrintTree(tree *Tree) {
	log.Println("pretty:", tree.Value)
	printTree(tree, "")
}
