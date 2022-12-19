package configparser

import (
	"os"
	"reflect"
	"testing"
)

type testConfig struct {
	IntValue    int    `json:"intValue" yaml:"intValue"`
	StringValue string `json:"stringValue" yaml:"stringValue"`
	BoolValue   bool   `json:"boolValue" yaml:"boolValue"`
	SliceValue  []int  `json:"sliceValue" yaml:"sliceValue"`
}

func TestParse(t *testing.T) {
	tests := []struct {
		name          string
		configType    ConfigFileType
		fileName      string
		target        interface{}
		defaults      map[string]interface{}
		envVars       map[string]string
		expectedError error
		expectedValue testConfig
	}{
		{
			name:          "non-pointer target",
			configType:    ConfigFileTypeJSON,
			fileName:      "",
			target:        testConfig{},
			defaults:      nil,
			envVars:       nil,
			expectedError: ErrTargetNotPointer,
			expectedValue: testConfig{},
		},
		{
			name:          "empty file name",
			configType:    ConfigFileTypeJSON,
			fileName:      "",
			target:        &testConfig{},
			defaults:      nil,
			envVars:       nil,
			expectedError: ErrEmptyFileName,
			expectedValue: testConfig{},
		},
		{
			name:          "non-existent file",
			configType:    ConfigFileTypeJSON,
			fileName:      "/path/to/non-existent/file",
			target:        &testConfig{},
			defaults:      nil,
			envVars:       nil,
			expectedError: ErrFileDoesNotExist,
			expectedValue: testConfig{},
		},
		{
			name:          "invalid config file type",
			configType:    ConfigFileType(255),
			fileName:      "./.fixture/config.json",
			target:        &testConfig{},
			defaults:      nil,
			envVars:       nil,
			expectedError: ErrInvalidConfigFileType,
			expectedValue: testConfig{},
		},
		{
			name:       "default values",
			configType: ConfigFileTypeJSON,
			fileName:   "./.fixture/incomplete-config.json",
			target:     &testConfig{},
			defaults: map[string]interface{}{
				"sliceValue": []int{1, 2, 3},
			},
			envVars:       nil,
			expectedError: nil,
			expectedValue: testConfig{
				IntValue:    123,
				StringValue: "test string",
				BoolValue:   true,
				SliceValue:  []int{1, 2, 3},
			},
		},
		{
			name:          "JSON test file variables",
			configType:    ConfigFileTypeJSON,
			fileName:      "./.fixture/config.json",
			target:        &testConfig{},
			defaults:      nil,
			envVars:       nil,
			expectedError: nil,
			expectedValue: testConfig{
				IntValue:    123,
				StringValue: "test string",
				BoolValue:   true,
				SliceValue:  []int{1, 2, 3},
			},
		},
		{
			name:          "YAML test file variables",
			configType:    ConfigFileTypeYAML,
			fileName:      "./.fixture/config.yaml",
			target:        &testConfig{},
			defaults:      nil,
			envVars:       nil,
			expectedError: nil,
			expectedValue: testConfig{
				IntValue:    123,
				StringValue: "test string",
				BoolValue:   true,
				SliceValue:  []int{1, 2, 3},
			},
		},
		{
			name:       "environment variables",
			configType: ConfigFileTypeJSON,
			fileName:   "./.fixture/config.json",
			target:     &testConfig{},
			defaults:   nil,
			envVars: map[string]string{
				"INTVALUE":    "789",
				"STRINGVALUE": "env string",
				"BOOLVALUE":   "false",
				"SLICEVALUE":  "4,5,6",
			},
			expectedError: nil,
			expectedValue: testConfig{
				IntValue:    789,
				StringValue: "env string",
				BoolValue:   false,
				SliceValue:  []int{4, 5, 6},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.envVars != nil {
				for k, v := range test.envVars {
					os.Setenv(k, v)
					defer os.Unsetenv(k)
				}
			}

			err := Parse(test.configType, test.fileName, test.target, test.defaults)
			if err != test.expectedError {
				t.Errorf("unexpected error: got %v, want %v", err, test.expectedError)
			}

			if test.expectedError == nil {
				if !reflect.DeepEqual(test.target, &test.expectedValue) {
					t.Errorf("unexpected value: got %v, want %v", test.target, test.expectedValue)
				}
			}

		})
	}
}
