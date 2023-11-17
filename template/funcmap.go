package template

import (
	"strings"
	"text/template"
)

// GenericFuncMap returns a generic template.FuncMap.
func GenericFuncMap() template.FuncMap {
	return funcMap
}

var (
	funcMap = map[string]any{
		"filterMap": filterMap,
		"splitN":    strings.SplitN,
		"mapLength": mapLength,
	}
)

// filterMap returns a map of string=any pairs that have the given key prefix.
// If the key prefix is empty, it returns the original map.
// If the keys in the original map do not have the given prefix, it returns an empty map.
func filterMap(data map[string]any, keyPrefix string) map[string]any {
	if keyPrefix == "" {
		return data
	}
	filtered := make(map[string]any)
	for key, value := range data {
		if strings.HasPrefix(key, keyPrefix) {
			filtered[key] = value
		}
	}
	return filtered
}

func mapLength(dt map[string]any) int {
	return len(dt)
}
