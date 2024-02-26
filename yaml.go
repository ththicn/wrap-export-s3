package main

import (
	"fmt"
	"os"

	"github.com/goccy/go-yaml"
)

type Resources struct {
	In struct {
		Type     string   `yaml:"type"`
		Database string   `yaml:"database"`
		Tables   []string `yaml:"tables"`
	} `yaml:"in"`
	Out struct {
		Type   string `yaml:"type"`
		Bucket string `yaml:"bucket"`
		Path   string `yaml:"path"`
	} `yaml:"out"`
}

// readFileContent reads the content of the file at the given path.
func readFileContent(filePath string) ([]byte, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("error reading YAML file: %s", err)
	}
	return content, nil
}

// unmarshalYAML unmarshals the YAML content into the Resources struct.
func unmarshalYAML(content []byte) (*Resources, error) {
	var res Resources
	err := yaml.Unmarshal(content, &res)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling YAML data: %s", err)
	}
	return &res, nil
}

// parseYAMLFile parses the YAML file at the given path and returns a Resources struct.
func parseYAMLFile(filePath string) (*Resources, error) {
	content, err := readFileContent(filePath)
	if err != nil {
		return nil, err
	}

	return unmarshalYAML(content)
}
