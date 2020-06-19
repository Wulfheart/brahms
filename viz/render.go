package viz

import (
	svg "github.com/ajstarks/svgo"
	"github.com/lucasb-eyer/go-colorful"
	"io"
	"wulfheart/brahms/score"
)

func Render (w io.Writer, scr score.Score) *svg.SVG{
	// // Score
	// maxPitch := scr.MaxPitch()
	// minPitch := scr.MinPitch()
	// maxDuration := scr.MaxDuration()
	// totalTickes := scr.TotalTicks()
	//
	// // Graphic
	// maxR := 200
	// minR := 50
	// maxRNode := 30
	//
	// colors, err := colorful.HappyPalette(len(scr))
	// if err != nil {
	// 	panic(err)
	// }
	// s := svg.New(w)
	// s.Start(500, 500)
	// for

	s.Circle(250, 250, 50, "stroke:rgb(255,0,0); fill: rgb(255,0,0); fill-opacity: 0.5")
	s.Circle(300, 240, 50, "fill: rgb(255,0,0); fill-opacity: 0.5")
	s.Circle(250, 250, 20, "fill: rgb(255,255,0); fill-opacity: 0.5")
	s.End()
	return s
}
