package configparser

import (
	"io/fs"
	"strings"

	"github.com/goccy/go-reflect"
	"github.com/spf13/viper"
)

// ConfigFileType is a type for constants that represent the type of a configuration file.
type ConfigFileType uint8

const (
	// ConfigFileTypeJSON represents a JSON configuration file.
	ConfigFileTypeJSON ConfigFileType = iota
	// ConfigFileTypeYAML represents a YAML configuration file.
	ConfigFileTypeYAML
)

var configFileTypeString = map[ConfigFileType]string{
	ConfigFileTypeJSON: "json",
	ConfigFileTypeYAML: "yaml",
}

// Parse reads in a configuration file and unmarshals the data into the provided target value.
// The configFileType argument specifies the type of the configuration file (JSON or YAML).
// The file argument is the file path of the configuration file. The target argument should be a pointer to the value
// that the configuration data will be unmarshaled into. The defaults argument is a map of default values that will be
// set in the configuration if they are not present in the configuration file.
// If there is an error reading in the configuration file or unmarshaling the data, an error is returned.
func Parse(configFileType ConfigFileType, file string, target interface{}, defaults map[string]interface{}) error {
	if reflect.ValueOf(target).Kind() != reflect.Ptr {
		return ErrTargetNotPointer
	}

	if file == "" {
		return ErrEmptyFileName
	}

	viper.Reset()

	configFileTypeString, ok := configFileTypeString[configFileType]
	if !ok {
		return ErrInvalidConfigFileType
	}

	viper.SetConfigType(configFileTypeString)

	for k, v := range defaults {
		viper.SetDefault(k, v)
	}

	viper.SetConfigFile(file)

	// https://github.com/spf13/viper/issues/584#issuecomment-451554896
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(*fs.PathError); ok && file != "" {
			return ErrFileDoesNotExist
		}

		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return err
		}
	}

	return viper.Unmarshal(target)
}
