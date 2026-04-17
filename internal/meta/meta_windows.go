//go:build windows

// Package meta provides file metadata such as hard link count, owner, and group.
// This is the Windows stub — Unix-specific fields are not available and return zero values.
package meta

import "os"

// Info holds file metadata. On Windows, all fields are empty/zero
// since Unix-specific syscall data is unavailable.
type Info struct {
	Links int    // always 0 on Windows
	Owner string // always empty on Windows
	Group string // always empty on Windows
}

// Get returns an empty Info on Windows since Unix-specific metadata is not available.
func Get(info os.FileInfo) Info {
	return Info{
		Links: 0,
		Owner: "",
		Group: "",
	}
}
