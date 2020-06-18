package score

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)







func Read(path string) Score{

	content, err := ioutil.ReadFile(path)
    if err != nil {
        log.Fatal(err)
    }
    // Convert []byte to string and print to screen
    csvfile := string(content)

    // This is some weird behaviour probably caused by one of the previous tools. Let's remove this, so it doesn't get weird and easily testable.
    csvfile = strings.ReplaceAll(csvfile, "\x00", "")
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
		for j, m := range l{
			l[j] = strings.ReplaceAll(m, "\r", "")
		}
		partNum := l[7]

		part, exists := Sc[partNum]
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
		part.Plays = append(part.Plays, &n)
		Sc[partNum] = part


	}
	return Sc
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
