package media

import (
	"fmt"
	"os"
	"path/filepath"
)

// FileRecord holds information about a media file
type FileRecord struct {
	FileName string // the actual name of the file (incl. extension)
	BaseName string // the base name of the file (excl. extension)
	Location string // the dir location of the file
	FileExt  string // the file extension
}

// NewFileRecordFromPath creates a new FileRecord from the provided `path`
func NewFileRecordFromPath(path string) *FileRecord {
	fileName := filepath.Base(path)
	ext := filepath.Ext(path)
	// trim extension
	name := fileName[:len(fileName)-len(ext)]
	return &FileRecord{
		FileName: fileName,
		BaseName: name,
		Location: filepath.Dir(path),
		FileExt:  ext,
	}
}

// NewFileRecordWithModeCheck checks if the file specified is a regular file (i.e. has no mode bits
// set for things like dir, symlink, temp file, etc.) before creating the FileRecord. If it is not a
// regular file, an error will be returned
func NewFileRecordWithModeCheck(path string, info os.FileInfo) (*FileRecord, error) {
	if info.Mode().IsRegular() {
		return NewFileRecordFromPath(path), nil
	} else {
		return nil, fmt.Errorf("Not a regular file")
	}
}
