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
	fmt.Println(len(colors))
	stepSize := float64(len(colors) - 1)/float64(steps - 1)
	for i := 0 ; i < steps; i++ {
		index := int(float64(i) * stepSize)
		if index > len(colors) - 1 {
			return nil, fmt.Errorf("calculated colorindex out of range")
		}
		pc = append(pc, colors[index])
	}
	return pc, nil
}

func MakeGradient(colors []colorful.Color) (rc []colorful.Color, err error) {
	for i := 0; i < len(colors)-1; i++ {
		steps := 1000
		for j := 1; j <= steps; j++ {
			rc = append(rc, colors[i].BlendHcl(colors[i+1], float64(j)/float64(steps - 1)).Clamped())
		}
	}
	return rc, nil
}
