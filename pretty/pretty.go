// Package pretty provides utilities for beautifying console output.
package pretty

import (
	"fmt"
	"log"
	"os"
	"strings"
)

// Progress represents a progress bar in the terminal.
type Progress struct {
	chardev bool
	current float64
}

// Update updates the progress bar to the specified percent (0 to 1).
func (p *Progress) Update(percent float64) {
	if percent > 1 {
		log.Panicln("pretty: the percent cannot be greater than 1")
	}
	if percent < p.current {
		log.Panicln("pretty: the percent cannot be decreased")
	}
	if percent != 0 && percent != 1 && percent-p.current < 0.01 {
		// Only update if the change is significant to avoid flickering.
		return
	}
	if percent == 1 && percent == p.current {
		// No need to update if already at 100%.
		return
	}
	if percent == 0 && p.chardev {
		// Save cursor position.
		log.Writer().Write([]byte{0x1b, 0x37})
	}
	if percent != 0 && p.chardev {
		// Load cursor position.
		log.Writer().Write([]byte{0x1b, 0x38})
	}
	p.current = percent
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

// NewProgress creates a new Progress instance.
func NewProgress() *Progress {
	s, err := os.Stdout.Stat()
	if err != nil {
		log.Panicln("pretty: cannot stat stdout:", err)
	}
	return &Progress{
		// Identify if we are displaying to a terminal or through a pipe or redirect.
		chardev: s.Mode()&os.ModeCharDevice == os.ModeCharDevice,
		current: 0,
	}
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
	Name string
	Leaf []*Tree
}

func (t *Tree) print(prefix string) {
	for i, elem := range t.Leaf {
		isLast := i == len(t.Leaf)-1
		branch := "├── "
		if isLast {
			branch = "└── "
		}
		log.Println("pretty:", prefix+branch+elem.Name)
		if len(elem.Leaf) > 0 {
			middle := "│   "
			if isLast {
				middle = "    "
			}
			elem.print(prefix + middle)
		}
	}
}

// PrintTree prints the tree structure starting from the root node.
func (t *Tree) Print() {
	log.Println("pretty:", t.Name)
	t.print("")
}

// NewTree creates a new Tree node with the given name.
func NewTree(name string) *Tree {
	return &Tree{Name: name}
}
