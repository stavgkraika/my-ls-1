package reader

import (
	"os"
	"path/filepath"
	"testing"

	"my-ls-1/internal/models"
)

// TestReadEntriesWithoutAFlagHidesDotfiles verifies that hidden files are excluded
// when the All flag is not set.
func TestReadEntriesWithoutAFlagHidesDotfiles(t *testing.T) {
	tmpDir := t.TempDir()

	createFile(t, filepath.Join(tmpDir, "visible.txt"))
	createFile(t, filepath.Join(tmpDir, ".hidden.txt"))

	entries, err := ReadEntries(tmpDir, models.Flags{All: false})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(entries) != 1 {
		t.Fatalf("expected 1 entry, got %d", len(entries))
	}

	if entries[0].Name != "visible.txt" {
		t.Fatalf("expected visible.txt, got %s", entries[0].Name)
	}
}

// TestReadEntriesWithAFlagIncludesDotfiles verifies that hidden files are included
// when the All flag is set.
func TestReadEntriesWithAFlagIncludesDotfiles(t *testing.T) {
	tmpDir := t.TempDir()

	createFile(t, filepath.Join(tmpDir, "visible.txt"))
	createFile(t, filepath.Join(tmpDir, ".hidden.txt"))

	entries, err := ReadEntries(tmpDir, models.Flags{All: true})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(entries) != 2 {
		t.Fatalf("expected 2 entries, got %d", len(entries))
	}
}

// TestJoinPath verifies that joinPath correctly joins a directory and file name.
func TestJoinPath(t *testing.T) {
	got := joinPath("/tmp", "file.txt")
	want := "/tmp/file.txt"

	if got != want {
		t.Fatalf("expected %s, got %s", want, got)
	}
}

// createFile is a test helper that creates an empty file at the given path.
func createFile(t *testing.T, path string) {
	t.Helper()

	file, err := os.Create(path)
	if err != nil {
		t.Fatalf("failed to create file %s: %v", path, err)
	}
	if err := file.Close(); err != nil {
		t.Fatalf("failed to close file %s: %v", path, err)
	}
}
