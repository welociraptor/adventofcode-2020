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
		if len(scanner.Text()) != 0 || !skipEmptyLines {
			data = append(data, scanner.Text())
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return data, nil
}
