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
	sc := score.Read(p)
	for _, v := range sc {
		v.Sort()
	}
	fmt.Println(sc.TotalTicks())
	fmt.Println("Sorting done")
}
