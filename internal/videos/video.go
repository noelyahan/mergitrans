package videos

import (
	"github.com/noelyahan/mergi"
	"image"
	"fmt"
	"github.com/noelyahan/mergi/io"
	"runtime"
	"path/filepath"
)


func getFrames(name string, from, to int) []image.Image {
	frames := make([]image.Image, 0)
	_, filename, _, _ := runtime.Caller(1)
	path := filepath.Join(filepath.Dir(filename), name)
	for i := from; i <= to; i++ {
		p := fmt.Sprintf("%s/frame%d.jpg", path, int(i))
		img, _ := mergi.Import(io.NewFileImporter(p))
		//img, _ = mergi.Resize(img, uint(img.Bounds().Max.X/2), uint(img.Bounds().Max.Y/2))
		//mergi.Export(io.NewFileExporter(img, fmt.Sprintf("%s1/img%d.jpg", path, int(i))))
		frames = append(frames, img)
	}
	return frames
}

var PoppyField = getFrames("1", 1, 150)
var Bird = getFrames("2", 1, 150)
var Clouds = getFrames("3", 1, 150)
