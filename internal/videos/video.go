package videos

import (
	"image"
	"fmt"
	"runtime"
	"path/filepath"
	"github.com/noelyahan/mergitrans/internal/io"
)

// GetFrames is used to load video frames
func GetFrames(name string, from, to int) []image.Image {
	frames := make([]image.Image, 0)
	_, filename, _, _ := runtime.Caller(1)
	path := filepath.Join(filepath.Dir(filename), "internal", "videos", name)
	for i := from; i <= to; i++ {
		p := fmt.Sprintf("%s/frame%d.jpg", path, int(i))
		img, _ := io.NewFileImporter(p).Import()
		frames = append(frames, img)
	}
	return frames
}