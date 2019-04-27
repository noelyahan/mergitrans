package internal_test

import (
	"testing"
	"github.com/noelyahan/mergi/transition"
	"github.com/noelyahan/mergitrans/internal/videos"
	"github.com/noelyahan/mergi"
	"github.com/noelyahan/mergi/io"
	"github.com/noelyahan/mergitrans"
)

func TestTransition(t *testing.T) {
	// transition init
	trans := mergitrans.ScaleDownFastCircle

	to := len(trans)

	// init videos
	video1 := videos.PoppyField
	video2 := videos.Clouds

	// splitted frame 1 - frame 5
	f1 := video1[0:to/2]
	f2 := video1[to/2:to+to/2]
	f3 := video2[0:to]
	f4 := video2[to:to+to/2]

	// transition will be applied to splitted frame 2 and frame 3
	imgs := transition.Images(
		f2,
		f3,
		trans, mergi.MaskBlack,
		0, float64(len(trans)-1), 1)

	// final frames will be combined and create a final video frames
	f5 := append(f1, imgs...)
	f5 = append(imgs, f4...)

	// create a animation using mergi lib
	anim, _:= mergi.Animate(f5, 1)
	mergi.Export(io.NewAnimationExporter(anim, "out.gif"))
}
