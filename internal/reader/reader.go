// Package reader handles reading directory entries from the filesystem.
package reader

import (
	"os"
	"strings"

	"my-ls-1/internal/models"
)

// ReadEntries reads the contents of the directory at path and returns a slice of Entry.
// Hidden files (starting with '.') are excluded unless the All flag is set.
func ReadEntries(path string, flags models.Flags) ([]models.Entry, error) {
	dirEntries, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	entries := make([]models.Entry, 0, len(dirEntries))

	for _, dirEntry := range dirEntries {
		name := dirEntry.Name()

		if shouldSkipHidden(name, flags) {
			continue
		}

		info, err := dirEntry.Info()
		if err != nil {
			continue
		}

		entries = append(entries, models.Entry{
			Name: name,
			Path: joinPath(path, name),
			Info: info,
		})
	}

	return entries, nil
}

// shouldSkipHidden reports whether a file should be excluded because it is hidden
// (i.e. its name starts with '.') and the All flag is not set.
func shouldSkipHidden(name string, flags models.Flags) bool {
	return !flags.All && strings.HasPrefix(name, ".")
}

// joinPath joins a base directory path and a file name into a full path,
// handling trailing slashes and the root "/" case correctly.
func joinPath(base, name string) string {
	if base == "/" {
		return "/" + name
	}
	if len(base) > 0 && base[len(base)-1] == '/' {
		return base + name
	}
	return base + "/" + name
}
