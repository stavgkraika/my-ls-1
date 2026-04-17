// Package sorter handles sorting of directory entries based on the provided flags.
package sorter

import (
	"sort"

	"my-ls-1/internal/models"
)

// SortEntries sorts the given entries in place according to the flags.
// It sorts by modification time if the Time flag is set, otherwise by name.
// If the Rev flag is set, the order is reversed after sorting.
func SortEntries(entries []models.Entry, flags models.Flags) {
	if flags.Time {
		sortByTime(entries)
	} else {
		sortByName(entries)
	}

	if flags.Rev {
		reverse(entries)
	}
}

// sortByName sorts entries alphabetically by name in ascending order.
func sortByName(entries []models.Entry) {
	sort.Slice(entries, func(i, j int) bool {
		return entries[i].Name < entries[j].Name
	})
}

// sortByTime sorts entries by modification time, newest first.
// Entries with equal modification times are sorted alphabetically by name as a tiebreaker.
func sortByTime(entries []models.Entry) {
	sort.Slice(entries, func(i, j int) bool {
		left := entries[i].Info.ModTime()
		right := entries[j].Info.ModTime()

		if left.Equal(right) {
			return entries[i].Name < entries[j].Name
		}

		return left.After(right)
	})
}

// reverse reverses the order of entries in place.
func reverse(entries []models.Entry) {
	for i, j := 0, len(entries)-1; i < j; i, j = i+1, j-1 {
		entries[i], entries[j] = entries[j], entries[i]
	}
}
