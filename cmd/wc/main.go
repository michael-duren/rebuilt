package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/michael-duren/rebuilt/internal/wc"
)

func main() {
	lines := flag.Bool("l", false, "count lines")
	words := flag.Bool("w", false, "count words")
	bytes := flag.Bool("c", false, "count bytes")
	chars := flag.Bool("m", false, "count characters, multi byte characters are included")
	longestLine := flag.Bool("L", false, "write the length of the line containing the most bytes")

	flag.Parse()

	files := flag.Args()

	opts := &wc.Options{
		CountLines:  *lines,
		CountWords:  *words,
		CountBytes:  *bytes,
		CountChars:  *chars,
		LongestLine: *longestLine,
	}

	output, err := wc.Run(files, opts)
	if err != nil {
		fmt.Fprintf(os.Stderr, "wc: %v\n", err)
		os.Exit(1)
	}
	fmt.Print(output)
}
