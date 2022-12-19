package configparser

import "errors"

var (
	// ErrTargetNotPointer is returned if the provided target value is not a pointer.
	ErrTargetNotPointer = errors.New("target is not pointer")
	// ErrEmptyFileName is returned if the provided file argument is an empty string.
	ErrEmptyFileName = errors.New("empty file name")
	// ErrFileDoesNotExist is returned if the provided file path does not exist.
	ErrFileDoesNotExist = errors.New("file does not exist")
	// ErrInvalidConfigFileType is returned if the provided configFileType constant is not a valid type.
	ErrInvalidConfigFileType = errors.New("invalid config file type")
)
