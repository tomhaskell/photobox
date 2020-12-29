package fs

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/tomhaskell/photobox/internal/pkg/media"
)

// Scan will search the provided path, and return a list of all the files
func Scan(path string) ([]*media.File, error) {
	files := make([]*media.File, 0)

	// find files
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		// loop over all items
		if err == nil {
			f := &media.File{
				FileName: filepath.Base(path),
				Location: filepath.Dir(path),
				FileType: filepath.Ext(path),
			}
			files = append(files, f)
		}
		return err
	})
	fmt.Println(err)

	return files, nil
}
