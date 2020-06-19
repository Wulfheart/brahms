package score

import (
	"fmt"
	"github.com/go-audio/midi"
	"os"
)

func ParseMidi(path string) Score{
	f, err := os.OpenFile(path, os.O_RDWR, os.ModeAppend)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	decoder := midi.NewDecoder(f)
	decoder.Debug = true
	err = decoder.Decode()

	t := make([]*midi.Event, 0)
	fmt.Println(len(decoder.Tracks[1].Events))
	for _, m := range decoder.Tracks[1].Events{
		// if m.Cmd == midi.EventByteMap["NoteOn"] || m.Cmd == midi.EventByteMap["NoteOff"] {
		// 	t = append(t, m)
		// }
		if m.Note != 0 {
			t = append(t, m)
		}
	}

	if err != nil {
		panic(err)
	}


	return nil
}