package source

import (
	"bytes"
	"os"

	"gopkg.in/yaml.v3"
)

func YAMLFile(path string) (map[string]any, error) {
	bts, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	out := map[string]any{}
	err = yaml.NewDecoder(bytes.NewBuffer(bts)).Decode(&out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
