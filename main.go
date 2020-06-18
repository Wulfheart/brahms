package main

import (
	"fmt"
	"path/filepath"
	"wulfheart/brahms/score"
)

func main() {
	p, err := filepath.Abs("./midi2csv/bach_brandenburg_3.csv")
	if err != nil {
		panic(err)
	}
	score.Read(p)
	fmt.Println(score.Score)
}
