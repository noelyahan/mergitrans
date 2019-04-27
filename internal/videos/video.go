package videos

import (
	"image"
	"fmt"
	"runtime"
	"path/filepath"
	"github.com/noelyahan/mergitrans/internal/io"
)


func getFrames(name string, from, to int) []image.Image {
	frames := make([]image.Image, 0)
	_, filename, _, _ := runtime.Caller(1)
	path := filepath.Join(filepath.Dir(filename), name)
	for i := from; i <= to; i++ {
		p := fmt.Sprintf("%s/frame%d.jpg", path, int(i))
		img, _ := io.NewFileImporter(p).Import()
		frames = append(frames, img)
	}
	return frames
}

var PoppyField = getFrames("1", 1, 150)
var Bird = getFrames("2", 1, 150)
var Clouds = getFrames("3", 1, 150)
