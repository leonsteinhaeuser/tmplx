package source

import (
	"encoding/json"
	"os"
	"path"
	"testing"
	"time"

	"gopkg.in/yaml.v3"
)

const (
	TestEnvKeyPrefix    = "GO_TEST_ENV_PREFIX_"
	TestEnvKeySubPrefix = "GO_TEST_ENV_PREFIX_SUB_"

	goTestFilesFolder = "go-test-files"
)

var (
	TestJSONFilePath = path.Join(os.TempDir(), goTestFilesFolder+time.Now().Format(time.RFC3339), "test.json")
	TestYAMLFilePath = path.Join(os.TempDir(), goTestFilesFolder+time.Now().Format(time.RFC3339), "test.yaml")

	envKeyValues = map[string]string{
		TestEnvKeyPrefix + "KEY1":    "value1",
		TestEnvKeyPrefix + "KEY2":    "value2",
		TestEnvKeySubPrefix + "KEY1": "value1",
		TestEnvKeySubPrefix + "KEY2": "value2",
	}
)

func TestMain(m *testing.M) {
	for k, v := range envKeyValues {
		err := os.Setenv(k, v)
		if err != nil {
			// ignore error
			continue
		}
	}

	// create test files folder
	_ = os.MkdirAll(path.Dir(TestJSONFilePath), 0755)

	jsonData, _ := json.Marshal(envKeyValues)
	yamlData, _ := yaml.Marshal(envKeyValues)

	// create test files
	_ = os.WriteFile(TestJSONFilePath, jsonData, 0644)
	_ = os.WriteFile(TestYAMLFilePath, yamlData, 0644)

	m.Run()

	for k := range envKeyValues {
		err := os.Unsetenv(k)
		if err != nil {
			// ignore error
			continue
		}
	}

	_ = os.RemoveAll(path.Join(os.TempDir(), goTestFilesFolder))
}
