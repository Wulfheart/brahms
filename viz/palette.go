package viz

import (
	"fmt"
	"github.com/lucasb-eyer/go-colorful"
)

func MakePalette(steps int, hex ...string) (pc []colorful.Color, err error) {
	var colors []colorful.Color
	for _, h := range hex {
		c, err := colorful.Hex(h)
		if err != nil {
			return nil, err
		}
		colors = append(colors, c)
	}
	colors, err = MakeGradient(colors)
	if err != nil {
		panic(err)
	}
	stepSize := float64(len(colors))/float64(steps)
	for i := 0 ; i < steps; i++ {
		index := int(float64(i) * stepSize)
		if index > len(colors) {
			return nil, fmt.Errorf("calculated colorindex out of range")
		}
		pc = append(pc, colors[index])
	}
	return pc, nil



	// return colors
}

func MakeGradient(colors []colorful.Color) (rc []colorful.Color, err error) {
	for i := 0; i < len(colors)-1; i++ {
		for j := 1; j < 100; j++ {
			rc = append(rc, colors[i].BlendHcl(colors[i+1], float64(i)/float64(j)))
		}
	}
	return rc, nil
}
