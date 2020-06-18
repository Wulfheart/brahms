package score

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"
)

var Score map[string]Part

type Part struct {
	Name string
	Plays []*Note
}

type Note struct {
	Part *Part
	StartTicks float64
	StartSecs float64
	DurTicks float64
	DurSecs float64
	Pitch float64
	FullNoteOctave float64
	Velocity float64
}

func Read(path string){
	csvfile, err := os.Open(path)
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}
	r := csv.NewReader(csvfile)
	r.FieldsPerRecord = -1
	out, err := r.ReadAll()
	if err != nil {
		panic(err)
	}
	for i, l := range out {
		if i == 0 {
			continue
		}
		if len(l) == 0 {
			continue
		}
		part, exists := Score[l[7]]
		if ! exists {
			part.Name = l[7]
		}
		n := Note{
			Part:           &part,
			StartTicks:     parseToFloat(l[0]),
			StartSecs:      parseToFloat(l[1]),
			DurTicks:       parseToFloat(l[2]),
			DurSecs:        parseToFloat(l[3]),
			Pitch:          parseToFloat(l[4]),
			FullNoteOctave: parseToFloat(l[5]),
			Velocity:       parseToFloat(l[6]),
		}
		part.Plays = append(part.Plays, &n)


	}
}

func parseToFloat(s string) float64{
	fmt.Println(reflect.TypeOf(s))
	v, err := strconv.ParseFloat(s, 64)
	if err != nil {
		panic(err)
	}
	return v
}
