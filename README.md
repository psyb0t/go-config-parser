# go-config-parser

This is a Go wrapper package for reading and unmarshalling data from configuration files based on [Viper](https://github.com/spf13/viper).

## Installation

To install this package, use `go get`:

```
go get github.com/psyb0t/go-config-parser
```

Import the package into your Go code:

```
import "github.com/psyb0t/go-config-parser"
```

## Usage

go-config-parser provides the `Parse` function for reading in and unmarshalling data from a configuration file into a target value. The function takes four arguments:

- configFileType: a constant representing the type of the configuration file (either `ConfigFileTypeJSON` or `ConfigFileTypeYAML`).
- file: the file path of the configuration file.
- target: a pointer to the value that the configuration data will be unmarshaled into.
- defaults: a map of default values that will be set in the configuration if they are not present in the configuration file.

The function uses the `viper` package to read in the configuration file, set the config file type, set default values, and unmarshal the data into the `target` value. It returns an error if there is a problem reading the configuration file or unmarshalling the data.

The function also handles setting environment variables as configuration sources using `viper.AutomaticEnv` to allow the configuration data to be overridden by environment variables. Using `viper.SetEnvKeyReplacer` it replaces periods in the environment variable keys with underscores to match the keys in the configuration file.

The order of importance for the used values are as such: environment variables > config file values > default values.

## Example

Here is an example of how to use the `Parse` function to read in a JSON configuration file and unmarshal the data into a struct:

```go
type Foo

type Config struct {
	IntValue    int    `json:"intValue"`
	StringValue string `json:"stringValue"`
	BoolValue   bool   `json:"boolValue"`
	SliceValue  []int  `json:"sliceValue"`
	Nested struct {
		Bar string `json:"bar"`
	} `json:"nested"`
}

var config Config

defaults := map[string]interface{}{
	"sliceValue": []int{1, 2, 3},
}

os.Setenv("NESTED_BAR", "test")

err := configparser.Parse(configparser.ConfigFileTypeJSON, "config.json", &config, defaults)
if err != nil {
	// handle error
}

// use config values
fmt.Println(config.IntValue)
fmt.Println(config.StringValue)
fmt.Println(config.BoolValue)
fmt.Println(config.SliceValue)
fmt.Println(config.Nested.Bar)
```

## Errors

The package defines several error constants that may be returned by the `Parse` function:

- `ErrTargetNotPointer`: returned if the provided `target` value is not a pointer.
- `ErrEmptyFileName`: returned if the provided `file` argument is an empty string.
- `ErrFileDoesNotExist`: returned if the provided file path does not exist.
- `ErrInvalidConfigFileType`: returned if the provided `configFileType` constant is not a valid type.
