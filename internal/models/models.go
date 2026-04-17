package models

import "os"

type Flags struct {
	Long bool // -l
	Rec  bool // -R
	All  bool // -a
	Rev  bool // -r
	Time bool // -t
}

type Entry struct {
	Name string
	Path string
	Info os.FileInfo
}
