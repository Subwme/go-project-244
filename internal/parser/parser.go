package parser

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

func Parse(filePath string) (map[string]interface{}, error) {
	absPath, err := filepath.Abs(filePath)
	if err != nil {
		return nil, err
	}

	data, err := os.ReadFile(absPath)
	if err != nil {
		return nil, err
	}

	switch filepath.Ext(filePath) {
	case ".json":
		return parseJSON(data)
	case ".yml", ".yaml":
		return parseYAML(data)
	default:
		return nil, fmt.Errorf("unsupported file format: %s", filepath.Ext(filePath))
	}
}

func parseJSON(data []byte) (map[string]interface{}, error) {
	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func parseYAML(data []byte) (map[string]interface{}, error) {
	var result map[string]interface{}
	if err := yaml.Unmarshal(data, &result); err != nil {
		return nil, err
	}
	return result, nil
}
