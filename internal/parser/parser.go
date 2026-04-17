// Package parser handles parsing of command-line arguments into flags and paths.
package parser

import (
	"errors"
	"strings"

	"my-ls-1/internal/models"
)

// ParseArgs parses the given command-line arguments into flags and a list of paths.
// If no paths are provided, it defaults to the current directory ".".
func ParseArgs(args []string) (models.Flags, []string, error) {
	var flags models.Flags
	var paths []string

	for _, arg := range args {
		if isFlag(arg) {
			if err := parseFlagGroup(arg, &flags); err != nil {
				return models.Flags{}, nil, err
			}
			continue
		}
		paths = append(paths, arg)
	}

	if len(paths) == 0 {
		paths = []string{"."}
	}

	return flags, paths, nil
}

// isFlag reports whether the argument is a flag (starts with '-' and has at least one character after it).
func isFlag(arg string) bool {
	return len(arg) > 1 && strings.HasPrefix(arg, "-")
}

// parseFlagGroup parses a single flag group (e.g. "-laR") and sets the corresponding fields on flags.
// Returns an error if an unknown flag character is encountered.
func parseFlagGroup(arg string, flags *models.Flags) error {
	if arg == "-" {
		return errors.New("invalid option -- '-'")
	}

	for _, ch := range arg[1:] {
		switch ch {
		case 'l':
			flags.Long = true
		case 'R':
			flags.Rec = true
		case 'a':
			flags.All = true
		case 'r':
			flags.Rev = true
		case 't':
			flags.Time = true
		default:
			return errors.New("invalid option -- '" + string(ch) + "'")
		}
	}

	return nil
}
