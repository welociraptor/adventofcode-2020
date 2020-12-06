package input

import (
	"bufio"
	"os"
)

func Strings(filename string, skipEmptyLines bool) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	data := make([]string, 0)

	for scanner.Scan() {
		row := scanner.Text()
		if len(row) != 0 || !skipEmptyLines {
			data = append(data, row)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return data, nil
}
