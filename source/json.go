package source

import (
	"bytes"
	"encoding/json"
	"os"
)

func JSONFile(path string) (map[string]any, error) {
	bts, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	out := map[string]any{}
	err = json.NewDecoder(bytes.NewBuffer(bts)).Decode(&out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
