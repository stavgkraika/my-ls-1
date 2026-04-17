package sorter

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	"my-ls-1/internal/models"
)

// TestSortEntriesByName verifies that entries are sorted alphabetically by name in ascending order.
func TestSortEntriesByName(t *testing.T) {
	tmpDir := t.TempDir()

	createFile(t, filepath.Join(tmpDir, "b.txt"))
	createFile(t, filepath.Join(tmpDir, "a.txt"))

	entries := mustReadTestEntries(t, tmpDir)
	SortEntries(entries, models.Flags{})

	if len(entries) != 2 {
		t.Fatalf("expected 2 entries, got %d", len(entries))
	}

	if entries[0].Name != "a.txt" || entries[1].Name != "b.txt" {
		t.Fatalf("unexpected order: %s, %s", entries[0].Name, entries[1].Name)
	}
}

// TestSortEntriesByNameReversed verifies that entries are sorted in reverse alphabetical order
// when the Rev flag is set.
func TestSortEntriesByNameReversed(t *testing.T) {
	tmpDir := t.TempDir()

	createFile(t, filepath.Join(tmpDir, "a.txt"))
	createFile(t, filepath.Join(tmpDir, "b.txt"))

	entries := mustReadTestEntries(t, tmpDir)
	SortEntries(entries, models.Flags{Rev: true})

	if entries[0].Name != "b.txt" || entries[1].Name != "a.txt" {
		t.Fatalf("unexpected reversed order: %s, %s", entries[0].Name, entries[1].Name)
	}
}

// TestSortEntriesByTime verifies that entries are sorted by modification time, newest first.
func TestSortEntriesByTime(t *testing.T) {
	tmpDir := t.TempDir()

	oldFile := filepath.Join(tmpDir, "old.txt")
	newFile := filepath.Join(tmpDir, "new.txt")

	createFile(t, oldFile)
	time.Sleep(1100 * time.Millisecond)
	createFile(t, newFile)

	entries := mustReadTestEntries(t, tmpDir)
	SortEntries(entries, models.Flags{Time: true})

	if entries[0].Name != "new.txt" || entries[1].Name != "old.txt" {
		t.Fatalf("unexpected time order: %s, %s", entries[0].Name, entries[1].Name)
	}
}

// TestSortEntriesByTimeReversed verifies that entries are sorted by modification time oldest first
// when both the Time and Rev flags are set.
func TestSortEntriesByTimeReversed(t *testing.T) {
	tmpDir := t.TempDir()

	oldFile := filepath.Join(tmpDir, "old.txt")
	newFile := filepath.Join(tmpDir, "new.txt")

	createFile(t, oldFile)
	time.Sleep(1100 * time.Millisecond)
	createFile(t, newFile)

	entries := mustReadTestEntries(t, tmpDir)
	SortEntries(entries, models.Flags{Time: true, Rev: true})

	if entries[0].Name != "old.txt" || entries[1].Name != "new.txt" {
		t.Fatalf("unexpected reversed time order: %s, %s", entries[0].Name, entries[1].Name)
	}
}

// mustReadTestEntries reads all entries from the given directory and returns them as a slice of Entry.
// It fails the test immediately if reading fails.
func mustReadTestEntries(t *testing.T, dir string) []models.Entry {
	t.Helper()

	dirEntries, err := os.ReadDir(dir)
	if err != nil {
		t.Fatalf("failed to read dir: %v", err)
	}

	entries := make([]models.Entry, 0, len(dirEntries))
	for _, de := range dirEntries {
		info, err := de.Info()
		if err != nil {
			t.Fatalf("failed to get info: %v", err)
		}

		entries = append(entries, models.Entry{
			Name: de.Name(),
			Path: filepath.Join(dir, de.Name()),
			Info: info,
		})
	}

	return entries
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
