package score

type Score map[string]*Part

func (s *Score) TotalTicks() int{
	maxTicks := 0
	for _, p := range *s {
		// Part Layer
		maxTicksPart := 0
		for _, n := range p.Plays{
			end := int(n.StartTicks + n.DurTicks)
			if end > maxTicksPart {
				maxTicksPart = end
			}
		}
		if maxTicksPart > maxTicks {
			maxTicks = maxTicksPart
		}
	}
	return maxTicks
}