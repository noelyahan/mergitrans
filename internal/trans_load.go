package internal

import (
	"image"
	"fmt"
	"runtime"
	"path/filepath"
	"github.com/noelyahan/mergitrans/internal/io"
)

// GetFrames used to load transition frames
func GetFrames(name string, from, to int) []image.Image {
	_, filename, _, _ := runtime.Caller(1)
	path := filepath.Join(filepath.Dir(filename), name)
	frames := make([]image.Image, 0)
	for i := from; i <= to; i++ {
		p := fmt.Sprintf("%s/frame%d.png", path, int(i))
		img, _ := io.NewFileImporter(p).Import()
		frames = append(frames, img)
	}
	return frames
}
