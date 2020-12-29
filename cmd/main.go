package main

import (
	"flag"
	"fmt"
	"strconv"

	"github.com/tomhaskell/photobox/internal/pkg/fs"
)

var (
	photoPath = flag.String("path", "~/Pictures/Photos/", "Source photo directory")
)

func main() {
	flag.Parse()

	fileList, err := fs.Scan(*photoPath)
	if err != nil {
		fmt.Println(err)
	}

	for _, imgFile := range fileList {
		fmt.Println(imgFile.FileName)
	}
	fmt.Println(strconv.Itoa(len(fileList)) + " files found")

}
