package score

import (
	"sort"
	"strconv"
)

type Score map[string]*Part

func (s *Score) SortedKeys() []string {
	keys := make([]string, 0)
	for k, _ := range *s {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		numA, _ := strconv.Atoi(keys[i])
		numB, _ := strconv.Atoi(keys[j])
		return numA < numB
	})
	return keys
}
func (s *Score) TotalTicks() float64 {
	maxTicks := 0.0
	for _, p := range *s {
		// Part Layer
		maxTicksPart := 0.0
		for _, n := range p.Plays {
			end := n.StartTicks + n.DurTicks
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
func (s *Score) MaxDuration() float64 {
	maxDurTicks := 0.0
	for _, p := range *s {
		// Part Layer
		maxDurationPart := 0.0
		for _, n := range p.Plays {
			end := n.DurTicks
			if end > maxDurationPart {
				maxDurationPart = end
			}
		}
		if maxDurationPart > maxDurTicks {
			maxDurTicks = maxDurationPart
		}
	}
	return maxDurTicks
}

func (s *Score) MinDuration() float64 {
	minDuration := 1000000.0
	for _, p := range *s {
		// Part Layer
		minDurationPart := 1000000.0
		for _, n := range p.Plays {
			end := n.Pitch
			if end < minDurationPart {
				minDurationPart = end
			}
		}
		if minDurationPart < minDuration {
			minDuration = minDurationPart
		}
	}
	return minDuration
}

func (s *Score) AvgDuration() float64 {
	totalDuration := 0.0
	i := 0
	for _, p := range *s {
		for _, n := range p.Plays {
			totalDuration += n.DurTicks
			i++
		}
	}
	return totalDuration / float64(i)
}

func (s *Score) MaxPitch() float64 {
	maxPitch := 0.0
	for _, p := range *s {
		// Part Layer
		maxPitchPart := 0.0
		for _, n := range p.Plays {
			end := n.Pitch
			if end > maxPitchPart {
				maxPitchPart = end
			}
		}
		if maxPitchPart > maxPitch {
			maxPitch = maxPitchPart
		}
	}
	return maxPitch
}

func (s *Score) MinPitch() float64 {
	// ! This is stupid. But 0 as minPitch works fine
	minPitch := 0.0
	for _, p := range *s {
		// Part Layer
		minPitchPart := 0.0
		for _, n := range p.Plays {
			end := n.Pitch
			if end < minPitchPart {
				minPitchPart = end
			}
		}
		if minPitchPart < minPitch {
			minPitch = minPitchPart
		}
	}
	return minPitch
}
