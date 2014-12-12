//
// Video

package main

import (
	"flag"
	"fmt"
	"os"
	"path"
	"path/filepath"
)

// MediaType My Custom Media Type
type MediaType int

// Media
// Media Type Constants
const (
	Video MediaType = 1 + iota // Video
	Image                      // Image
	Sound                      // Sound
)

func (m MediaType) String() string {
	switch m {
	case Video:
		return "Video"
	case Image:
		return "Image"
	case Sound:
		return "Sound"
	}
	return ""
}

// Media Represents a media file on the file system
// TODO
type Media struct {
	Path string
	Type MediaType
}

func (m Media) String() string {
	return fmt.Sprintf("%s@%s", m.Type, m.Path)
}

// ScanDir Scan a directory
func ScanDir(dirPath string) []Media {

	mediaMap := map[string]MediaType{
		".mp4": Video,
		".jpg": Image,
		".png": Image,
	}

	items := []Media{}

	var walkFunc = func(root string, info os.FileInfo, err error) error {
		ext := mediaMap[path.Ext(info.Name())]
		if ext == Video {
			items = append(items, Media{root, ext})
		}
		return nil
	}

	filepath.Walk(dirPath, walkFunc)

	return items
}

func main() {
	flag.Parse()
	dirPath := flag.Arg(0)
	for _, m := range ScanDir(dirPath) {
		fmt.Println(m)
	}
}
