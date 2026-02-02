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
