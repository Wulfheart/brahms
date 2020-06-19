package main

import (
	"fmt"
	"path/filepath"
	"wulfheart/brahms/score/midicsv"
)

func main() {
	p, err := filepath.Abs("./midi2csv/bach_brandenburg_3.mid")
	if err != nil {
		panic(err)
	}
	st, err := midicsv.Process(p)
	if err != nil {
		panic(err)
	}
	fmt.Println(st)

}
