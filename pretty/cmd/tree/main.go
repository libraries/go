package main

import (
	"log"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/libraries/go/pretty"
)

func walk(path string) *pretty.Tree {
	info, err := os.Stat(path)
	if err != nil {
		log.Panicln("main:", err)
	}
	node := &pretty.Tree{Value: info.Name()}
	if info.IsDir() {
		l, err := os.ReadDir(path)
		if err != nil {
			log.Panicln("main:", err)
		}
		for _, e := range l {
			node.Elems = append(node.Elems, walk(filepath.Join(path, e.Name())))
		}
		// Sort the elements alphabetically for consistent output.
		slices.SortFunc(node.Elems, func(a, b *pretty.Tree) int {
			return strings.Compare(a.Value, b.Value)
		})
	}
	return node
}

func main() {
	pretty.PrintTree(walk("."))
}
