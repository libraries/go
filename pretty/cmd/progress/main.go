package main

import (
	"time"

	"github.com/libraries/go/pretty"
)

func main() {
	progress := pretty.NewProgress()
	progress.Update(0)
	for i := range 1024 {
		time.Sleep(time.Millisecond * 4)
		progress.Update(float64(i+1) / 1024)
	}
	progress.Update(1)
}
