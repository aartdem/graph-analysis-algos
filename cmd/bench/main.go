package main

import (
	"fmt"
	"graph-analysis-algos/graphs"
	"os"
	"path/filepath"
	"time"
)

var (
	files, _ = filepath.Glob("data/*.mtx")
	iters    = 20
)

func main() {
	for _, path := range files {
		g, err := graphs.LoadMTX(path)
		if err != nil {
			fmt.Fprintf(os.Stderr, "read %s: %v\n", path, err)
			return
		}

		base := filepath.Base(path)

		fmt.Printf("%s", base)
		// warm up
		for i := 0; i < 2; i++ {
			_ = graphs.EdgeConnectivity(g)
		}
		for i := 0; i < iters; i++ {
			start := time.Now()
			_ = graphs.EdgeConnectivitySW(g) // change algo here
			el := time.Since(start)
			fmt.Printf(",%.0f", float64(el.Microseconds())/1000.0)
		}
		fmt.Println()
	}
	return
}
