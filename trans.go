package mergitrans

import (
	"github.com/noelyahan/mergitrans/internal"
	"github.com/noelyahan/mergitrans/internal/videos"
	"image"
)

// bonus video frames
var Videos = struct {
	PoppyField []image.Image
	Bird []image.Image
	Clouds []image.Image
}{videos.PoppyField, videos.Bird, videos.Clouds}

// Support transitions
var SlideBar = internal.GetSlideBar()
var Ink1 = internal.GetInk1()
var Ink2 = internal.GetInk2()
var Ink3 = internal.GetInk3()
var ScaleUpFastRect = internal.GetScaleUpFastRect()
var ScaleDownFastRect = internal.GetScaleDownFastRect()
var ScaleUpFastCircle = internal.GetScaleUpFastCircle()
var ScaleDownFastCircle = internal.GetScaleDownFastCircle()
