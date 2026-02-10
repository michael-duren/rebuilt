package sorting

import (
	"bufio"
	"os"
)

func Run(files []string) (string, error) {
	return "", nil
}

func readFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if scanner.Err() != nil {
		return lines, err
	}
	return lines, nil
}
