// Package core is the entry point for listing files.
// It coordinates reading and sorting entries for a given path.
package core

import (
	"my-ls-1/internal/models"
	"my-ls-1/internal/reader"
	"my-ls-1/internal/sorter"
)

// GetFiles reads and returns sorted directory entries for the given path,
// applying the provided flags to filter and order the results.
func GetFiles(path string, flags models.Flags) ([]models.Entry, error) {
	entries, err := reader.ReadEntries(path, flags)
	if err != nil {
		return nil, err
	}

	sorter.SortEntries(entries, flags)
	return entries, nil
}
