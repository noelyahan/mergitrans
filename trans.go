package mergitrans

import (
	"github.com/noelyahan/mergitrans/internal"
	"github.com/noelyahan/mergitrans/internal/videos"
	"image"
)

// bonus video frames
var Videos = struct {
	PoppyField func() []image.Image
	Bird       func() []image.Image
	Clouds     func() []image.Image
}{
	videos.GetFrames("1", 1, 150),
	videos.GetFrames("2", 1, 150),
	videos.GetFrames("3", 1, 150)}

// support transitions
var SlideBar = internal.GetFrames("slide_bar", 1, 36)
var Ink1 = internal.GetFrames("transition_ink", 1, 60)
var Ink2 = internal.GetFrames("transition_ink", 61, 98)
var Ink3 = internal.GetFrames("transition_ink", 99, 150)
var ScaleUpFastRect = internal.GetFrames("scale_rect", 39, 69)
var ScaleDownFastRect = internal.GetFrames("scale_rect", 1, 39)
var ScaleUpFastCircle = internal.GetFrames("scale_circle", 39, 69)
var ScaleDownFastCircle = internal.GetFrames("scale_circle", 1, 39)
