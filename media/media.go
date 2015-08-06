//
// Video

package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"time"
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
		".avi": Video,
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
	rand.Seed(time.Now().UnixNano())
	flag.Parse()

	dirPath := flag.Arg(0)

	items := ScanDir(dirPath)
	pick := items[rand.Intn(len(items))]

	fmt.Println(pick)
	exec.Command("open", pick.Path).Run()
}
