package main

import (
	"time"

	"github.com/libraries/go/pretty"
)

func main() {
	pretty.PrintProgress(0)
	for i := range 100 {
		time.Sleep(time.Millisecond * 16)
		pretty.PrintProgress(float64(i+1) / 100)
	}
	pretty.PrintProgress(1)
}
