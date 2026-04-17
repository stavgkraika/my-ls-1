package parser

import "testing"

// TestParseArgsFlagsOnly verifies that combined flags are parsed correctly
// and that the default path "." is used when no path is given.
func TestParseArgsFlagsOnly(t *testing.T) {
	flags, paths, err := ParseArgs([]string{"-laR"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !flags.Long || !flags.All || !flags.Rec {
		t.Fatalf("flags not parsed correctly: %+v", flags)
	}

	if len(paths) != 1 || paths[0] != "." {
		t.Fatalf("expected default path '.', got %v", paths)
	}
}

// TestParseArgsFlagsAndPath verifies that flags and an explicit path are both parsed correctly.
func TestParseArgsFlagsAndPath(t *testing.T) {
	flags, paths, err := ParseArgs([]string{"-tr", "testdir"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !flags.Time || !flags.Rev {
		t.Fatalf("flags not parsed correctly: %+v", flags)
	}

	if len(paths) != 1 || paths[0] != "testdir" {
		t.Fatalf("expected path testdir, got %v", paths)
	}
}

// TestParseArgsMultiplePaths verifies that multiple path arguments are all captured.
func TestParseArgsMultiplePaths(t *testing.T) {
	_, paths, err := ParseArgs([]string{"dir1", "dir2"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(paths) != 2 || paths[0] != "dir1" || paths[1] != "dir2" {
		t.Fatalf("unexpected paths: %v", paths)
	}
}

// TestParseArgsInvalidFlag verifies that an unknown flag character returns an error.
func TestParseArgsInvalidFlag(t *testing.T) {
	_, _, err := ParseArgs([]string{"-z"})
	if err == nil {
		t.Fatal("expected error for invalid flag")
	}
}
