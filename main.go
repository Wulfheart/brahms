package main

import (
	"fmt"
	"github.com/lucasb-eyer/go-colorful"
	"path/filepath"
	"wulfheart/brahms/score"
	"wulfheart/brahms/viz"
)

func main() {
	p, err := filepath.Abs("./midi2csv/bach_brandenburg_3.csv")
	if err != nil {
		panic(err)
	}
	sc := score.Read(p)
	for _, v := range sc {
		v.Sort()
	}
	viz.Test(sc)
	colors, err := colorful.HappyPalette(11)
	if err != nil {
		panic(err)
	}
	for _, c := range colors {
		fmt.Println(c.Hex())
	}
}
