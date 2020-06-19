package viz

import (
	"fmt"
	svg "github.com/ajstarks/svgo"
	"github.com/lucasb-eyer/go-colorful"
	"io"
	"math"
	"wulfheart/brahms/score"
)

func RenderCircle(w io.Writer, scr score.Score) *svg.SVG {
	// Score
	maxPitch := scr.MaxPitch()
	minPitch := scr.MinPitch()
	// maxDuration := scr.MaxDuration()
	avgDuration := scr.AvgDuration()
	totalTicks := scr.TotalTicks()

	// Graphic
	maxR := 350.0
	minR := 0.0
	maxRNode := 15.0

	wd := 800
	h := 800

	c1, _ := colorful.Hex("#3CA55C")
	c2, _ := colorful.Hex("#B5AC49")
	colors := makePalette(c1, c2, len(scr))
	s := svg.New(w)
	s.Start(wd, h)
	for i, k := range scr.SortedKeys() {
		p := scr[k]
		color := colors[i]
		for _, n := range p.Plays {
			r := ((n.Pitch)/(maxPitch-minPitch))*(maxR-minR) + minR
			if r == 0 {
				// panic(fmt.Sprintf("%f,%f,%f,%f,%f", n.Pitch, maxPitch, minPitch, maxR, minR))
			}
			// TODO: Plus or minus (mathematically positive?)
			angle := (n.StartTicks / totalTicks) * 2 * math.Pi
			if angle > 2*math.Pi {
				panic(angle)
			}
			// nodeR := math.Ceil((n.DurTicks / maxDuration) * maxRNode)
			nodeR := (n.DurTicks/avgDuration) * 1/4 * maxRNode
			if nodeR > maxRNode {
				nodeR = maxRNode
			}
			x, y := cartesian(r, angle)
			style := fmt.Sprintf("stroke: %s; fill: %s; fill-opacity: 0.3", "none", color.Hex())
			s.Circle(wd/2+int(x), h/2+int(y), int(nodeR), style)
		}

	}
	s.End()
	return s
}

func makePalette(c1 colorful.Color, c2 colorful.Color, steps int) []colorful.Color {
	var colors []colorful.Color
	for i := 0; i < steps; i++ {
		colors = append(colors, c1.BlendHcl(c2, float64(i)/float64(steps-1)))
	}
	return colors
}
