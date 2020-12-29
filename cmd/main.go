package main

import (
	"context"
	"flag"
	"fmt"
	"strconv"

	"github.com/tomhaskell/photobox/internal/pkg/fs"
)

var (
	photoPath = flag.String("path", "~/Pictures/Photos/", "Source photo directory")
	threads   = flag.Int("threads", 10, "Number of processing pipelines")
)

func main() {
	flag.Parse()

	ctx := context.Background()

	fileList, err := fs.Scan(*photoPath, *threads, ctx)
	if err != nil {
		fmt.Println(err)
	}

	for _, imgFile := range fileList {
		fmt.Println(imgFile.FileName)
	}
	fmt.Println(strconv.Itoa(len(fileList)) + " files found")

}
