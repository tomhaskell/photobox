package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/tomhaskell/photobox/internal/pkg/media"
)

var (
	imgPath = flag.String("img", "", "Source photo path")
	asJSON  = flag.Bool("json", false, "Print tags as json")
)

func main() {
	flag.Parse()

	entries, err := media.FindExif(*imgPath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if *asJSON {
		data, err := json.MarshalIndent(entries, "", "    ")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(data))
	} else {
		for _, tag := range entries {
			fmt.Printf("Name: %s, Value: %s\n", tag.TagName, tag.Formatted)
		}
	}

}
