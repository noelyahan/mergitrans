package videos

import (
	"image"
	"fmt"
	"runtime"
	"path/filepath"
	"github.com/noelyahan/eximp"
)

// GetFrames is used to load video frames
func GetFrames(name string, from, to int) func() []image.Image {
	frames := make([]image.Image, 0)
	_, filename, _, _ := runtime.Caller(1)
	path := filepath.Join(filepath.Dir(filename), "internal", "videos", name)
	return func() []image.Image {
		p := ""
		for i := from; i <= to; i++ {
			p = fmt.Sprintf("%s/frame%d.jpg", path, int(i))
			img, _ := eximp.NewFileImporter(p).Import()
			frames = append(frames, img)
		}
		return frames
	}
}