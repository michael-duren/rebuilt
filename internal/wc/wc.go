package wc

import "io"

type Options struct {
	CountLines  bool
	CountWords  bool
	CountBytes  bool
	CountChars  bool
	LongestLine bool
}

func Run(files []string, opts Options) error {
	return nil
}

func processInput(r io.Reader, opts Options) error {
	return nil
}
