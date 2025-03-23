package parser

import (
	"bufio"
	"os"
	"strings"
)

func LoadEnv(path string) (map[string]interface{}, error) {
	result := make(map[string]interface{})
	if path == "" {
		return result, nil
	}
	file, err := os.Open(path)
	if err != nil {
		return result, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			result[parts[0]] = strings.TrimSpace(parts[1])
		}
	}
	return result, scanner.Err()
}
