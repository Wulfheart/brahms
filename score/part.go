package score

import "sort"

type Part struct {
	Name string
	Plays []*Note
}

func (p *Part) Sort(){
	sort.Slice(p.Plays, func(i, j int) bool {
		return p.Plays[i].StartTicks < p.Plays[j].StartTicks
		})
}
