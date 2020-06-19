package main

import (
	"fmt"
	"path/filepath"
	"wulfheart/brahms/score"
	"wulfheart/brahms/viz"
)

func main() {
	fmt.Println("Compiling finished")
	p, err := filepath.Abs("./midi2csv/bach_brandenburg_3.mid")
	if err != nil {
		panic(err)
	}
	sc := score.Read(p)
	fmt.Println(sc["1"].Plays[0].DurSecs)
	viz.Test(sc)
	// st, err := midicsv.Process(p)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Print(st)

}
