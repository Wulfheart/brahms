package score

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