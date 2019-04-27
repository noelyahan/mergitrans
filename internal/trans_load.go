package internal

import (
	"image"
	"fmt"
	"runtime"
	"path/filepath"
	"github.com/noelyahan/mergitrans/internal/io"
)

func GetSlideBar() []image.Image {
	return getFrames("slide_bar", 1, 36)
}

func GetInk1() []image.Image {
	return getFrames("transition_ink", 1, 60)
}

func GetInk2() []image.Image {
	return getFrames("transition_ink", 61, 98)
}

func GetInk3() []image.Image {
	return getFrames("transition_ink", 99, 150)
}

func GetScaleUpFastRect() []image.Image {
	return getFrames("scale_rect", 39, 69)
}

func GetScaleDownFastRect() []image.Image {
	return getFrames("scale_rect", 1, 39)
}

func GetScaleUpFastCircle() []image.Image {
	return getFrames("scale_circle", 39, 69)
}

func GetScaleDownFastCircle() []image.Image {
	return getFrames("scale_circle", 1, 39)
}

func getFrames(name string, from, to int) []image.Image {
	_, filename, _, _ := runtime.Caller(1)
	path := filepath.Join(filepath.Dir(filename), "../", name)
	frames := make([]image.Image, 0)
	for i := from; i <= to; i++ {
		p := fmt.Sprintf("%s/frame%d.png", path, int(i))
		img, _ := io.NewFileImporter(p).Import()
		frames = append(frames, img)
	}
	return frames
}
