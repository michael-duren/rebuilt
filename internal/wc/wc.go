package wc

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type Options struct {
	CountLines  bool
	CountWords  bool
	CountBytes  bool
	CountChars  bool
	LongestLine bool
}

func (o *Options) NonePassed() bool {
	return !o.CountLines &&
		!o.CountWords &&
		!o.CountBytes &&
		!o.CountChars &&
		!o.LongestLine
}

type wordCount struct {
	filename    string
	lines       int64
	words       int64
	bytes       int64
	chars       int64
	longestLine int64
}

func Run(files []string, opts *Options) (string, error) {
	wordCounts := []wordCount{}
	for _, filename := range files {
		wc, err := processFile(filename)
		if err != nil {
			return "", err
		}

		wordCounts = append(wordCounts, *wc)
	}

	return formatResult(wordCounts, opts), nil
}

func formatResult(wordCounts []wordCount, opts *Options) string {
	sb := strings.Builder{}

	for _, wc := range wordCounts {
		sb.WriteString(fmt.Sprintf("%s ", wc.filename))
		if opts.NonePassed() {
			sb.WriteString(fmt.Sprintf("word: %d, line: %d, character: %d, byte: %d\n", wc.words, wc.lines, wc.chars, wc.bytes))
			continue
		}
		if opts.CountLines {
			sb.WriteString(fmt.Sprintf("lines: %d ", wc.lines))
		}
		if opts.CountWords {
			sb.WriteString(fmt.Sprintf("words: %d ", wc.words))
		}
		if opts.CountBytes {
			sb.WriteString(fmt.Sprintf("bytes: %d ", wc.bytes))
		}
		if opts.CountChars {
			sb.WriteString(fmt.Sprintf("chars: %d ", wc.chars))
		}
		if opts.LongestLine {
			sb.WriteString(fmt.Sprintf("longest line: %d ", wc.longestLine))
		}
		sb.WriteByte('\n')
	}
	return sb.String()
}

func NewWordCount(filename string, lines, words, bytes, countChars, longestLine int64) *wordCount {
	return &wordCount{
		filename, lines, words, bytes, countChars, longestLine,
	}
}

func (wc *wordCount) AddLine(lc *lineCounts) {
	wc.lines++
	wc.bytes += lc.bytes
	wc.chars += lc.chars
	wc.words += lc.words
	wc.longestLine = max(wc.longestLine, lc.bytes)
}

func processFile(filename string) (wc *wordCount, fileErr error) {
	file, err := os.Open(filename)

	if err != nil {
		return nil, err
	}

	defer func() {
		closeErr := file.Close()
		if closeErr != nil && fileErr == nil {
			fileErr = err
		}
	}()
	return process(file, filename)
}

func process(r io.Reader, filename string) (wc *wordCount, fileErr error) {

	wc = NewWordCount(filename, 0, 0, 0, 0, 0)

	scanner := bufio.NewReader(r)
	for {
		line, err := scanner.ReadString('\n')
		if err == io.EOF {
			if len(line) > 0 {
				lc := countLine(line)
				wc.AddLine(lc)
			}

			break
		}

		if err != nil {
			return nil, err
		}
		lc := countLine(line)
		wc.AddLine(lc)
	}

	return wc, nil
}

type lineCounts struct {
	words int64
	bytes int64
	chars int64
}

func countLine(line string) *lineCounts {
	lc := &lineCounts{}
	lc.bytes = int64(len(line))
	lc.chars = int64(len([]rune(line)))
	lc.words = int64(len(strings.Fields(line)))

	return lc
}
