package wc

import (
	"strings"
	"testing"
)

func TestProcess(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		wantLines   int64
		wantWords   int64
		wantBytes   int64
		wantChars   int64
		wantLongest int64
	}{
		{
			name:        "basic two lines",
			input:       "hello world\nfoo bar baz\n",
			wantLines:   2,
			wantWords:   5,
			wantBytes:   24,
			wantChars:   24,
			wantLongest: 12,
		},
		{
			name:        "empty input",
			input:       "",
			wantLines:   0,
			wantWords:   0,
			wantBytes:   0,
			wantChars:   0,
			wantLongest: 0,
		},
		{
			name:        "no trailing newline",
			input:       "hello world",
			wantLines:   1,
			wantWords:   2,
			wantBytes:   11,
			wantChars:   11,
			wantLongest: 11,
		},
		{
			name:        "multibyte characters",
			input:       "héllo 世界\n",
			wantLines:   1,
			wantWords:   2,
			wantBytes:   14,
			wantChars:   9,
			wantLongest: 14,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := strings.NewReader(tt.input)
			wc, err := process(input, "test")
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if wc.lines != tt.wantLines {
				t.Errorf("lines = %d, want %d", wc.lines, tt.wantLines)
			}
			if wc.words != tt.wantWords {
				t.Errorf("words = %d, want %d", wc.words, tt.wantWords)
			}
			if wc.bytes != tt.wantBytes {
				t.Errorf("bytes = %d, want %d", wc.bytes, tt.wantBytes)
			}
			if wc.chars != tt.wantChars {
				t.Errorf("chars = %d, want %d", wc.chars, tt.wantChars)
			}
			if wc.longestLine != tt.wantLongest {
				t.Errorf("longestLine = %d, want %d", wc.longestLine, tt.wantLongest)
			}
		})
	}
}

func TestFormatResult(t *testing.T) {
	tests := []struct {
		name       string
		wordCounts []wordCount
		opts       *Options
		want       string
	}{
		{
			name: "no options passed shows all",
			wordCounts: []wordCount{
				{filename: "test.txt", lines: 10, words: 50, bytes: 200, chars: 180, longestLine: 40},
			},
			opts: &Options{},
			want: "test.txt word: 50, line: 10, character: 180, byte: 200\n",
		},
		{
			name: "lines only",
			wordCounts: []wordCount{
				{filename: "test.txt", lines: 10, words: 50, bytes: 200, chars: 180, longestLine: 40},
			},
			opts: &Options{CountLines: true},
			want: "test.txt lines: 10 \n",
		},
		{
			name: "words only",
			wordCounts: []wordCount{
				{filename: "test.txt", lines: 10, words: 50, bytes: 200, chars: 180, longestLine: 40},
			},
			opts: &Options{CountWords: true},
			want: "test.txt words: 50 \n",
		},
		{
			name: "bytes only",
			wordCounts: []wordCount{
				{filename: "test.txt", lines: 10, words: 50, bytes: 200, chars: 180, longestLine: 40},
			},
			opts: &Options{CountBytes: true},
			want: "test.txt bytes: 200 \n",
		},
		{
			name: "chars only",
			wordCounts: []wordCount{
				{filename: "test.txt", lines: 10, words: 50, bytes: 200, chars: 180, longestLine: 40},
			},
			opts: &Options{CountChars: true},
			want: "test.txt chars: 180 \n",
		},
		{
			name: "longest line only",
			wordCounts: []wordCount{
				{filename: "test.txt", lines: 10, words: 50, bytes: 200, chars: 180, longestLine: 40},
			},
			opts: &Options{LongestLine: true},
			want: "test.txt longest line: 40 \n",
		},
		{
			name: "multiple options",
			wordCounts: []wordCount{
				{filename: "test.txt", lines: 10, words: 50, bytes: 200, chars: 180, longestLine: 40},
			},
			opts: &Options{CountLines: true, CountWords: true},
			want: "test.txt lines: 10 words: 50 \n",
		},
		{
			name: "multiple files",
			wordCounts: []wordCount{
				{filename: "a.txt", lines: 5, words: 20, bytes: 100, chars: 90, longestLine: 25},
				{filename: "b.txt", lines: 3, words: 10, bytes: 50, chars: 45, longestLine: 20},
			},
			opts: &Options{CountLines: true},
			want: "a.txt lines: 5 \nb.txt lines: 3 \n",
		},
		{
			name:       "empty wordCounts",
			wordCounts: []wordCount{},
			opts:       &Options{},
			want:       "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := formatResult(tt.wordCounts, tt.opts)
			if got != tt.want {
				t.Errorf("formatResult() = %q, want %q", got, tt.want)
			}
		})
	}
}
