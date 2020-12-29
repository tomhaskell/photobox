package fs

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github.com/tomhaskell/photobox/internal/pkg/media"
)

// Scan will search the provided path, and return a list of all the files
func Scan(path string, concurrency int, ctx context.Context) ([]*media.File, error) {
	files := make([]*media.File, 0)

	// concurrency sync
	var wg sync.WaitGroup

	// create pipeline streams
	pathStream := make(chan string)
	fileStream := make(chan *media.File)

	// setup workers
	for x := 0; x < concurrency; x++ {
		go makeFileWorker(pathStream, fileStream, &wg)
	}
	wg.Add(concurrency)

	// find files
	go func() {
		err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
			// loop over all items
			if err == nil {
				pathStream <- path
			}
			return err
		})
		fmt.Println(err)
		close(pathStream)
	}()

	// process files
	go func() {
		for f := range fileStream {
			files = append(files, f)
		}
	}()

	wg.Wait()
	close(fileStream)

	return files, nil
}

func makeFileWorker(pathStream chan string, fileStream chan *media.File, wg *sync.WaitGroup) {
	for path := range pathStream {
		//fmt.Printf("Processing [%v]\n", path)
		fileStream <- &media.File{
			FileName: filepath.Base(path),
			Location: filepath.Dir(path),
			FileType: filepath.Ext(path),
		}
	}
	fmt.Println("ended")
	wg.Done()
}
