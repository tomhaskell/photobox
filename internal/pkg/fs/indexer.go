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
func Scan(ctx context.Context, path string, concurrency int) ([]*media.FileRecord, error) {
	files := make([]*media.FileRecord, 0)

	// concurrency sync
	var wg sync.WaitGroup
	wg.Add(concurrency)

	// create pipeline streams
	fileStream := make(chan *media.FileRecord)

	// setup workers
	for x := 0; x < concurrency; x++ {
		go processFileWorker(ctx, &wg, fileStream)
	}

	// find files
	go startStream(ctx, fileStream, path)

	// process files
	go func() {
		for f := range fileStream {
			files = append(files, f)
		}
	}()

	wg.Wait()

	return files, nil
}

func startStream(ctx context.Context, fileStream chan *media.FileRecord, path string) error {
	// find all files in path and it's subdirs
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		//fmt.Printf("Processing [%v]\n", path)
		if err == nil {
			select {
			case <-ctx.Done():
				return ctx.Err()
			default:
				f, err2 := media.NewFileRecordWithModeCheck(path, info)
				if err2 == nil {
					fileStream <- f
				}
			}
		}
		return err
	})
	close(fileStream)
	return err
}

func processFileWorker(ctx context.Context, wg *sync.WaitGroup, fileStream chan *media.FileRecord) {
	defer wg.Done()
	for f := range fileStream {
		fmt.Printf("Processing [%v]\n", f.FileName)
		//select {
		//case fileStream <- media.NewFileRecordWithModeCheck(path):
		//case <-ctx.Done():
		// return
	}

	// }
	fmt.Println("ended")
}
