package input

import (
	"bufio"
	"os"
	"strconv"
)

func ReadFileToString(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

func ReadFileToStringSlice(path string) ([]string, error) {
	ret := make([]string, 0)
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ret = append(ret, scanner.Text())
	}

	return ret, nil
}

func ReadFileToIntSlice(path string) ([]int, error) {
	data, err := ReadFileToStringSlice(path)
	if err != nil {
		return nil, err
	}

	ret := make([]int, len(data))
	for i, v := range data {
		ret[i], err = strconv.Atoi(v)
		if err != nil {
			return nil, err
		}
	}

	return ret, nil
}
