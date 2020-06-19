package score

import (
	"sort"
	"strconv"
)

type Score map[string]*Part

func (s *Score) SortedKeys() []string{
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
func (s *Score) TotalTicks() int {
	maxTicks := 0
	for _, p := range *s {
		// Part Layer
		maxTicksPart := 0
		for _, n := range p.Plays {
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
func (s *Score) MaxDuration() int {
	maxDurTicks := 0
	for _, p := range *s {
		// Part Layer
		maxDurationPart := 0
		for _, n := range p.Plays {
			end := int(n.DurTicks)
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
func (s *Score) MaxPitch() int {
	maxPitch := 0
	for _, p := range *s {
		// Part Layer
		maxPitchPart := 0
		for _, n := range p.Plays {
			end := int(n.Pitch)
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

func (s *Score) MinPitch() int {
	minPitch := 0
	for _, p := range *s {
		// Part Layer
		minPitchPart := 0
		for _, n := range p.Plays {
			end := int(n.Pitch)
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
