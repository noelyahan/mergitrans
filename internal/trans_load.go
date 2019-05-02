package internal

import (
	"image"
	"fmt"
	"runtime"
	"path/filepath"
	"github.com/noelyahan/eximp"
)

// GetFrames used to load transition frames
func GetFrames(name string, from, to int) func() []image.Image {
	_, filename, _, _ := runtime.Caller(1)
	path := filepath.Join(filepath.Dir(filename), name)
	frames := make([]image.Image, 0)
	return func() []image.Image {
		p := ""
		var img image.Image
		for i := from; i <= to; i++ {
			p = fmt.Sprintf("%s/frame%d.png", path, int(i))
			//fmt.Print(p)
			img, _ = eximp.NewFileImporter(p).Import()
			frames = append(frames, img)
		}
		return frames
	}
}
