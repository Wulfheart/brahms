package score

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"log"
	"os/exec"
	"sort"
	"strconv"
	"strings"
)

// Requires midicsv installed
func Midi2csv(p string) {
	cmd := exec.Command("midicsv", p)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	csvfile := out.String()
	csvfile = strings.ReplaceAll(csvfile, "\x00", "")
	r := csv.NewReader(strings.NewReader(csvfile))
	r.FieldsPerRecord = -1
	r.LazyQuotes = true
	read, err := r.ReadAll()
	if err != nil {
		panic(err)
	}
	fmt.Println(read[0])

	// 	Filter
	for o, m := range read {
		for p, n := range m {
			read[o][p] = strings.TrimSpace(n)
		}
	}
	deleter := make([]int, 0)
	for i, l := range read {
		if !lineNeeded(l[2]) {
			deleter = append(deleter, i)
		}
	}
	sort.Ints(deleter)
	deleter = reverseInts(deleter)
	for _, d := range deleter {
		read = append(read[:d], read[d+1:]...)
	}

	// 	Process data
	type tempo struct {
		track  string
		tick   int
		number int
	}

	var tempi []tempo

	// var header struct {
	// 	format   int
	// 	division int
	// 	nTracks  int
	// }
	for _, row := range read {
		switch row[2] {
		case "Header":
			break
		case "Note_on_c":
			break
		case "Note_off_c":
			break
		case "Tempo":
			tick, err := strconv.Atoi(row[1])
			if err != nil {
				panic(err)
			}
			number, err := strconv.Atoi(row[3])
			if err != nil {
				panic(err)
			}
			tempi = append(tempi, tempo{
				track:  row[0],
				tick:   tick,
				number: number,
			})
			break
		}
	}
	fmt.Println("HERE")
}

func lineNeeded(third string) bool {
	return third == "Note_on_c" || third == "Note_off_c" || third == "Tempo" || third == "Header"

}

func reverseInts(a []int) []int {
	for left, right := 0, len(a)-1; left < right; left, right = left+1, right-1 {
		a[left], a[right] = a[right], a[left]
	}
	return a
}
