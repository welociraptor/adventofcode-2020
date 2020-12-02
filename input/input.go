package input

import (
	"bufio"
	"os"
)

func Strings(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	data := make([]string, 0)

	for scanner.Scan() {
		t := scanner.Text()
		if len(t) > 0 {
			data = append(data, scanner.Text())
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return data, nil
}
