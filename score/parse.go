package score

import (
	"encoding/csv"
	"fmt"
	"github.com/Wulfheart/brahms/score/midicsv"
	"io/ioutil"
	"strconv"
	"strings"
)

const (
	ReadProcessedCsv = iota
	ReadMidi
)

func Read(path string, opt int, midicsv_location string) Score {
	var csvfile string
	if opt == ReadMidi {
		o, err := midicsv.Process(path, midicsv_location)
		if err != nil {
			panic(err)
		}
		csvfile = o
	} else if opt == ReadProcessedCsv {
		f, err := ioutil.ReadFile(path)
		if err != nil {
			panic(err)
		}
		csvfile = string(f)
	} else {
		panic("No valid option.")
	}
	csvfile = strings.ReplaceAll(csvfile, "\x00", "")

	// This is some weird behaviour probably caused by one of the previous tools. Let's remove this, so it doesn't get weird and easily testable.
	r := csv.NewReader(strings.NewReader(csvfile))
	r.FieldsPerRecord = -1
	out, err := r.ReadAll()
	if err != nil {
		panic(err)
	}
	Sc := make(Score)
	for i, l := range out {
		if i == 0 {
			continue
		}
		if len(l) != 8 {
			continue
		}
		if len(l) != 8 {
			fmt.Println(i)
		}
		for j, m := range l {
			l[j] = strings.ReplaceAll(m, "\r", "")
		}
		partNum := l[7]

		part, exists := Sc[partNum]
		if !exists {
			part = &Part{
				Name:  l[7],
				Plays: nil,
			}
		}

		n := Note{
			Part:           part,
			StartTicks:     parseToFloat(l[0]),
			StartSecs:      parseToFloat(l[1]) / 1000000.0,
			DurTicks:       parseToFloat(l[2]),
			DurSecs:        parseToFloat(l[3]) / 1000000.0,
			Pitch:          parseToFloat(l[4]),
			FullNoteOctave: l[5],
			Velocity:       parseToFloat(l[6]),
		}
		part.Plays = append(part.Plays, &n)
		Sc[partNum] = part

	}
	return Sc
}

func parseToFloat(s string) float64 {
	v, err := strconv.ParseFloat(s, 64)
	if err != nil {
		panic(err)
	}
	return v
}
