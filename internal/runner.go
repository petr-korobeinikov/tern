package internal

import (
	"errors"
	"strings"
)

func Run(args []string) (string, error) {
	var arg, truthy, falsy string

	switch len(args) {
	case shellMode:
		arg, truthy, falsy = args[0], args[1], args[2]

		if len(arg) > 0 {
			return truthy, nil
		}
		return falsy, nil
	case boolMode:
		if args[0] != "-b" {
			return "", ErrUnsupportedMode
		}

		arg, truthy, falsy = args[1], args[2], args[3]

		if arg == "0" || strings.ToLower(arg) == "false" || len(arg) == 0 {
			return falsy, nil
		}
		return truthy, nil
	default:
		return "", ErrIncorrectNumberOfArguments
	}
}

var (
	ErrUnsupportedMode            = errors.New("unsupported mode")
	ErrIncorrectNumberOfArguments = errors.New("incorrect number of arguments")
)

const (
	shellMode = 3
	boolMode  = 4
)
