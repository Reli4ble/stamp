package parser

import (
	"os"

	"gopkg.in/yaml.v3"
)

func LoadYAML(path string) (map[string]interface{}, error) {
	data := make(map[string]interface{})
	if path == "" {
		return data, nil
	}
	b, err := os.ReadFile(path)
	if err != nil {
		return data, err
	}
	err = yaml.Unmarshal(b, &data)
	return data, err
}

func MergeMaps(base, override map[string]interface{}) map[string]interface{} {
	out := make(map[string]interface{})
	for k, v := range base {
		out[k] = v
	}
	for k, v := range override {
		out[k] = v
	}
	return out
}
