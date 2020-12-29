package media

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/dsoprea/go-exif/v3"
)

func FindExif(path string) ([]exif.ExifTag, error) {
	// open and read file
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	data, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	// find start of exif data
	rawExif, err := exif.SearchAndExtractExif(data)
	if err != nil {
		if err == exif.ErrNoExif {
			fmt.Printf("No EXIF data for %v.\n", path)
		}
		return nil, err
	}

	fmt.Printf("[DEBUG] EXIF blob is (%d) bytes.", len(rawExif))

	// parse the exif data into ExifTags
	entries, _, err := exif.GetFlatExifData(rawExif, nil)
	if err != nil {
		return nil, err
	}

	return entries, nil

}
