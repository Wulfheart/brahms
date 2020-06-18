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
	val, ok := score.Score["1"]
	fmt.Println("Value:", val, ok)
	for k, i := range score.Score {
		fmt.Println(k, i.Name)
	}
}
