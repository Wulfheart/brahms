package score

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

var Score = make(map[string]*Part)

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
	FullNoteOctave string
	Velocity float64
}

func Read(path string){
	  content, err := ioutil.ReadFile(path)
    if err != nil {
        log.Fatal(err)
    }

    // Convert []byte to string and print to screen
    csvfile := string(content)
	r := csv.NewReader(strings.NewReader(csvfile))
	r.FieldsPerRecord = -1
	out, err := r.ReadAll()
	if err != nil {
		panic(err)
	}
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
		for j, m := range l{
			l[j] = strings.ReplaceAll(m, "\x00", "")
		}
		part, exists := Score[l[7]]
		if ! exists {
			part = &Part{
				Name:  l[7],
				Plays: nil,
			}
		}

		n := Note{
			Part:           part,
			StartTicks:     parseToFloat(l[0]),
			StartSecs:      parseToFloat(l[1]),
			DurTicks:       parseToFloat(l[2]),
			DurSecs:        parseToFloat(l[3]),
			Pitch:          parseToFloat(l[4]),
			FullNoteOctave: l[5],
			Velocity:       parseToFloat(l[6]),
		}
		// fmt.Println(n)
		part.Plays = append(part.Plays, &n)
		fmt.Println(l[7])
		Score[l[7]] = part


	}
}

func parseToFloat(s string) float64{
	// Why is this needed? I have no idea
	s = strings.ReplaceAll(s, "\x00", "")
	v, err := strconv.ParseFloat(s, 64)
	if err != nil {
		panic(err)
	}
	return v
}
