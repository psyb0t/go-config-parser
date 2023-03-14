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
// The file argument is the file path of the configuration file.
// The target argument should be a pointer to the value that the configuration data will be unmarshaled into.
// The defaults argument is a map of default values that will be set in the
// configuration if they are not present in the configuration file.
// The envPrefix argument is a string specifying the environment variables
// prefix to use when reading env vars (eg. "myPrefix" results in "MYPREFIX_MYPROP")
// If there is an error reading in the configuration file or unmarshaling the data, an error is returned.
func Parse(configFileType ConfigFileType, file string, target interface{},
	defaults map[string]interface{}, envPrefix string) error {
	vpr := viper.New()

	if envPrefix != "" {
		vpr.SetEnvPrefix(envPrefix)
	}

	if reflect.ValueOf(target).Kind() != reflect.Ptr {
		return ErrTargetNotPointer
	}

	configFileTypeString, ok := configFileTypeString[configFileType]
	if !ok {
		return ErrInvalidConfigFileType
	}

	vpr.SetConfigType(configFileTypeString)

	for k, v := range defaults {
		vpr.SetDefault(k, v)
	}

	vpr.AutomaticEnv()
	vpr.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if file != "" {
		vpr.SetConfigFile(file)
		if err := vpr.ReadInConfig(); err != nil {
			if _, ok := err.(*fs.PathError); ok && file != "" {
				return ErrFileDoesNotExist
			}

			if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
				return err
			}
		}
	}

	return vpr.Unmarshal(target)
}
