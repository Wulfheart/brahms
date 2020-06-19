package viz

import (
	"bytes"
	svg "github.com/ajstarks/svgo"
	"wulfheart/brahms/score"
)

var s score.Score

func CreateCircular(score score.Score, buffer *bytes.Buffer) (s *svg.SVG) {
	// var buf *bytes.Buffer
	s = RenderCircle(buffer, score)
	return s
}

func CreateArcs(score score.Score, buffer *bytes.Buffer) (s *svg.SVG) {
	return RenderArc(buffer, score)
}
