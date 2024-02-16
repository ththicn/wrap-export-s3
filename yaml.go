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

// yamlファイルをparseする関数
func parseYAMLFile(filePath string) (*Resources, error) {
	// Read file content
	content, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("error reading YAML file: %s", err)
	}

	// Unmarshal the YAML into the Config struct
	var res Resources
	err = yaml.Unmarshal(content, &res)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling YAML data: %s", err)
	}

	// resの値のポインタを毎回作っているので呼び出すたびに別のポインタになる
	// つまり別アドレス
	return &res, nil
}
