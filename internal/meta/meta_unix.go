//go:build linux || darwin

// Package meta provides file metadata such as hard link count, owner, and group.
package meta

import (
	"os"
	"os/user"
	"strconv"
	"syscall"
)

// Info holds Unix-specific file metadata extracted from the syscall layer.
type Info struct {
	Links int    // number of hard links
	Owner string // name of the file owner
	Group string // name of the file group
}

// Get extracts Unix-specific metadata from the given os.FileInfo.
// It resolves numeric UID/GID to human-readable names when possible.
func Get(info os.FileInfo) Info {
	stat, ok := info.Sys().(*syscall.Stat_t)
	if !ok {
		return Info{}
	}

	owner := strconv.FormatUint(uint64(stat.Uid), 10)
	group := strconv.FormatUint(uint64(stat.Gid), 10)

	if u, err := user.LookupId(owner); err == nil {
		owner = u.Username
	}

	if g, err := user.LookupGroupId(group); err == nil {
		group = g.Name
	}

	return Info{
		Links: int(stat.Nlink),
		Owner: owner,
		Group: group,
	}
}
