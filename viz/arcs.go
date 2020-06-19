package viz

import (
	"fmt"
	svg "github.com/ajstarks/svgo"
	"github.com/lucasb-eyer/go-colorful"
	"io"
	"math"
	"wulfheart/brahms/score"
)

func RenderArc(w io.Writer, scr score.Score) *svg.SVG {
	maxPitch := scr.MaxPitch()
	minPitch := scr.MinPitch()
	maxDuration := scr.MaxDuration()
	totalTicks := scr.TotalTicks()

	// Graphic
	maxR := 350.0
	minR := 0.0

	c1, _ := colorful.Hex("#3CA55C")
	c2, _ := colorful.Hex("#B5AC49")
	colors := makePalette(c1, c2, len(scr))

	wd := 800
	h := 800

	s := svg.New(w)
	s.Start(wd, h)
	for i, k := range scr.SortedKeys() {
		p := scr[k]
		color := colors[i]
		for _, n := range p.Plays {
			r := ((n.Pitch)/(maxPitch-minPitch))*(maxR-minR) + minR
			angleStart := (n.StartTicks / totalTicks) * 2 * math.Pi
			angleEnd := ((n.StartTicks + n.DurTicks) / totalTicks) * 2 * math.Pi
			xS, yS := cartesian(r, angleStart)
			xE, yE := cartesian(r, angleEnd)

			s.Path(formatArc(int(xS)+wd/2, int(yS)+h/2, int(xE)+wd/2, int(yE)+h/2, int(r)), fmt.Sprintf("fill:none; stroke: %s; stroke-opacity: 0.5; stroke-width:%f;", color.Hex(), (n.DurTicks/maxDuration)*15))
		}
	}
	// s.Path(formatArc(200, 200, 300, 300, 20), "fill:none;stroke:red")
	s.End()

	return s
}

func formatArc(x1 int, y1 int, x2 int, y2 int, r int) string {
	return fmt.Sprintf("M %d %d A %d %d 0 0 0  %d %d", x1, y1, r, r, x2, y2)
}
