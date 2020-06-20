package viz

import (
	"fmt"
	svg "github.com/ajstarks/svgo"
	"github.com/lucasb-eyer/go-colorful"
	"io"
	"math"
	"wulfheart/brahms/score"
)

type CircleConfig struct {
	MaxR     float64
	MinR     float64
	MaxRNode float64
	Width    float64
	Height   float64
	Colors []colorful.Color
	Style    string
}

func RenderCircle(w io.Writer, scr score.Score, opt CircleConfig) *svg.SVG {
	// Score
	maxPitch := scr.MaxPitch()
	minPitch := scr.MinPitch()
	// maxDuration := scr.MaxDuration()
	avgDuration := scr.AvgDuration()
	totalTicks := scr.TotalTicks()

	// Graphic
	maxR := opt.MaxR
	minR := opt.MinR
	maxRNode := opt.MaxRNode

	wd := int(opt.Width)
	h := int(opt.Height)

	s := svg.New(w)
	s.Start(wd, h)
	for i, k := range scr.SortedKeys() {
		p := scr[k]
		color := opt.Colors[i]
		for _, n := range p.Plays {
			r := ((n.Pitch)/(maxPitch-minPitch))*(maxR-minR) + minR
			if r == 0 {
				// panic(fmt.Sprintf("%f,%f,%f,%f,%f", n.Pitch, maxPitch, minPitch, maxR, minR))
			}
			// TODO: Plus or minus (mathematically positive?)
			angle := (n.StartTicks/totalTicks)*2*math.Pi - 0.5*math.Pi
			if angle > 2*math.Pi {
				panic(angle)
			}
			// nodeR := math.Ceil((n.DurTicks / maxDuration) * maxRNode)
			nodeR := (n.DurTicks / avgDuration) * 1 / 5 * maxRNode
			if nodeR > maxRNode {
				nodeR = maxRNode
			}
			x, y := cartesian(r, angle)
			style := fmt.Sprintf("stroke: %s; fill: %s; fill-opacity: 0.5", "none", color.Hex())
			s.Circle(wd/2+int(x), h/2+int(y), int(nodeR), style)
		}

	}
	s.End()
	return s
}
