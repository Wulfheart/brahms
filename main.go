package main

import (
	"path/filepath"
	"wulfheart/brahms/score"
)

func main() {
	p, err := filepath.Abs("./midi2csv/bach_brandenburg_3.mid")
	if err != nil {
		panic(err)
	}
	score.Midi2csv(p)

}
