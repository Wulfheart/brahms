package midicsv

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"os/exec"
	"sort"
	"strconv"
	"strings"
)

// Rewritten in Go from https://github.com/shoogle/midicsv-process/blob/master/midicsv-process.py

var noteLetters = []string{"C", "C", "D", "D", "E", "F", "F", "G", "G", "A", "A", "B"}
var sharps = []string{"", "#", "", "#", "", "", "#", "", "#", "", "#", ""}

type noteEvent struct {
	track    int
	tick     int
	pitch    int
	velocity int
}

func initNoteEvent(track int, tick int, pitch int, velocity int) *noteEvent {
	return &noteEvent{
		track:    track,
		tick:     tick,
		pitch:    pitch,
		velocity: velocity,
	}
}

type note struct {
	track    int
	tick     int
	pitch    int
	velocity int
	duration int
}

func initNote(on noteEvent, off noteEvent) *note {
	return &note{
		track:    on.track,
		tick:     on.tick,
		pitch:    on.pitch,
		velocity: on.velocity,
		duration: off.tick - on.tick,
	}
}

func (n note) onTimeMicros(tm tempoMap) int {
	return tm.microsAtTick(n.tick)
}

func (n note) durationMicros(tm tempoMap) int {
	return tm.microsAtTick(n.tick+n.duration) - n.onTimeMicros(tm)
}
func (n note) octave() int {
	// TODO: Checking. Is this correct?
	return (n.pitch / 12) - 1
}

func (n note) letter() string {
	return noteLetters[n.pitch % 12]
}

func (n note) sharp() string {
	return sharps[n.pitch % 12]
}

func (n note) fullNote() string {
	return fmt.Sprintf("%s%s", n.letter(), n.sharp())
}

func (n note) fullNoteOctave() string {
	return fmt.Sprintf("%s%d%s", n.letter(), n.octave(), n.sharp())
}

func (n note) toString(tm tempoMap) string{
	return fmt.Sprintf("%d,%d,%d,%d,%d,%s,%d,%d", n.tick, n.onTimeMicros(tm), n.duration, n.durationMicros(tm), n.pitch, n.fullNoteOctave(), n.velocity, n.track)
}

type tempoEvent struct {
	tick   int
	micros int
	tempo  int
}

func initTempoEvent(tick int, micros int, tempo int) *tempoEvent {
	if tempo == 0 {
		tempo = 500000
	}
	return &tempoEvent{
		tick:   tick,
		micros: micros,
		tempo:  tempo,
	}
}

type tempoMap struct {
	tpqn int // Ticks per quarter note
	tmap []tempoEvent
}

func initTempoMap(tpqn int, tmap []tempoEvent) *tempoMap{
	if tpqn == 0 {
		tpqn = 480
	}
	return &tempoMap{
		tpqn: tpqn,
		tmap: tmap,
	}
}

func (t *tempoMap) addTempo(tick int, tempo int) error {
	t.tmap = append(t.tmap, *initTempoEvent(tick, tempo, 0))
	return nil
}

func (t tempoMap) tempoEventAtTick(tick int) tempoEvent {
	savedTempoEvent := *initTempoEvent(0, 0, 0)
	for _, te := range t.tmap {
		if te.tick > tick {
			break
		}
		savedTempoEvent = te
	}
	return savedTempoEvent
}

func (t tempoMap) microsAtTick(tick int) int {
	te := t.tempoEventAtTick(tick)
	return te.micros + ((tick - te.tick) * te.tempo / t.tpqn)
}

func Process(path string) (string, error){
	tempoMap := initTempoMap(0, nil)
	notes := make([]note, 0)
	noteEvents := make([]noteEvent, 0)
	onTicks := make([]int, 0)

	// Reading csv
	conv, err := convertMidi2Csv(path)
	if err != nil {
		return "", err
	}
	// Read MIDI Events
	for _, row := range conv {
		track, err := strconv.Atoi(row[0])
		if err != nil {
			return "", nil
		}
		tick, err := strconv.Atoi(row[1])
		if err != nil {
			return "", nil
		}
		typ := row[2]

		switch typ {
		case "Header":
			tpqn, err := strconv.Atoi(row[5])
			if err != nil {
				return "", nil
			}
			tempoMap.tpqn = tpqn
			break
		case "Tempo":
			tempo, err := strconv.Atoi(row[3])
			if err != nil {
				return "", nil
			}
			err = tempoMap.addTempo(tick, tempo)
			if err != nil {
				return "", nil
			}
			break
		case "Note_on_c":
			pitch, err := strconv.Atoi(row[4])
			if err != nil {
				return "", nil
			}
			velocity, err := strconv.Atoi(row[5])
			if err != nil {
				return "", nil
			}
			noteEvents = append(noteEvents, *initNoteEvent(track, tick, pitch, velocity))
			break
		}
	}
	// create notes by pairing noteOn and noteOff events
	for i, neon := range noteEvents {
		if neon.velocity == 0 {
			continue
		}
		for _, neoff := range noteEvents[i:] {
			if neoff.velocity != 0 || neoff.track != neon.track || neoff.pitch != neon.pitch {
				continue
			}
			note := initNote(neon, neoff)
			notes = append(notes, *note)
			onTicks = append(onTicks, note.tick)
			break
		}
	}

	// Sort notes by onTick
	sort.Slice(notes, func(i, j int) bool {
		return notes[i].tick < notes[j].tick
	})
	// generate csv string
	var b bytes.Buffer
	_, err = b.WriteString("start_ticks,start_secs,dur_ticks,dur_secs,pitch,fullNoteOctave,velocity,part\n")
	if err != nil {
		return "", err
	}
	for _, n := range notes {
		_, err = b.WriteString(n.toString(*tempoMap) + "\n")
		if err != nil {
			return "", err
		}
	}
	return b.String(), nil
}

func convertMidi2Csv(path string)([][]string, error){
	cmd := exec.Command("midicsv", path)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return nil, err
	}
	csvfile := out.String()
	csvfile = strings.ReplaceAll(csvfile, "\x00", "")
	r := csv.NewReader(strings.NewReader(csvfile))
	r.FieldsPerRecord = -1
	r.LazyQuotes = true
	read, err := r.ReadAll()
	if err != nil {
		return nil, err
	}

	// 	Filter
	for o, m := range read {
		for p, n := range m {
			read[o][p] = strings.TrimSpace(n)
		}
	}
	return read, nil

}


