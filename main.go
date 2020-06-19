package main

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"wulfheart/brahms/score"
	"wulfheart/brahms/viz"
)

func main() {
	fmt.Println("Compiling finished")
	// p, err := filepath.Abs("./midi2csv/Beethoven_9.mid")
	p, err := filepath.Abs("./midi2csv/bach_brandenburg_3.mid")
	// p, err := filepath.Abs("./midi2csv/pal.csv")
	if err != nil {
		panic(err)
	}
	sc := score.Read(p, score.ReadMidi)
	fmt.Println(sc.AvgDuration(), sc.MaxDuration(), sc.MinDuration())
	buf := new(bytes.Buffer)
	viz.CreateCircular(sc, buf)
	p, err = filepath.Abs("./midi2csv/out.svg")
	if err != nil {
		panic(err)
	}
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	_, err = f.Write(buf.Bytes())
	if err != nil {
		panic(err)
	}

}
