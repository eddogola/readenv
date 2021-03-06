package readenv

import (
	"bufio"
	"io"
	"os"
	"regexp"
)

// EnvData maps environment keys to values.
type EnvData map[string]string

// Get returns the value of given key from the map,
// returns error if key not found
func (data EnvData) Get(key string) (string, error) {
	val, ok := data[key]

	if !ok {
		return "", errKeyNotInEnvData(key)
	}
	return val, nil
}

func (data EnvData) add(key, val string) {
	data[key] = val
}

// Read takes in an io.Reader, reads env vars from it and returns an EnvData
func Read(reader io.Reader) (EnvData, error) {
	envData := make(EnvData)

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		key, val, err := parse(line)
		if err != nil {
			continue
		}
		envData.add(key, val)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return envData, nil
}

// File takes in a filename, reads env vars from it and returns an EnvData
func File(filename string) (EnvData, error) {
	file, err := os.OpenFile(filename, os.O_RDONLY, 0444) // open .env file
	if err != nil {
		return nil, err
	}
	defer file.Close()

	envData, err := Read(file)
	if err != nil {
		return nil, err
	}

	return envData, nil
}

// utility function to get key and value from line(string) provided
func parse(line string) (string, string, error) {
	re, err := regexp.Compile(`(\w+[^=])(?:=)([a-zA-Z0-9_-]+)`)
	if err != nil {
		return "", "", err
	}
	matches := re.FindStringSubmatch(line)
	if len(matches) != 3 {
		return "", "", errFindingTwoRegexSubmatches(line)
	}

	key := matches[1]
	val := matches[2]

	return key, val, nil
}
