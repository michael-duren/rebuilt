package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/michael-duren/rebuilt/internal/sorting"
)

func main() {
	files := flag.Args()
	res, err := sorting.Run(files)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error sorting %q: %v", files, err)
		os.Exit(1)
	}
	fmt.Fprintf(os.Stdout, res)
}
