package source

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

var (
	ErrEnvKeyNotInValidFormat = errors.New("env key is not in valid format")
)

// Environ returns a map of the environment variables with the given prefix.
func Environ(prefix string) (map[string]any, error) {
	envMap := map[string]any{}
	for _, env := range os.Environ() {
		vals := strings.Split(env, "=")
		if !strings.HasPrefix(vals[0], prefix) {
			// skip if not match prefix or not key=value format
			continue
		}
		if len(vals) != 2 {
			return nil, fmt.Errorf("%w: %s", ErrEnvKeyNotInValidFormat, env)
		}
		envMap[vals[0]] = vals[1]
	}
	return envMap, nil
}
