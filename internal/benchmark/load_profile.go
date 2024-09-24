package benchmark

import (
	"time"
)

type LoadProfile struct {
	Type     string
	BaseLoad int
	MaxLoad  int
	Duration time.Duration
}

func GenerateLoadProfile(profile LoadProfile) []int {
	var load []int
	duration := int(profile.Duration.Seconds())

	switch profile.Type {
	case "fixed":
		for i := 0; i < duration; i++ {
			load = append(load, profile.MaxLoad)
		}
	case "ramp_up":
		rampUpDuration := duration / 2
		for i := 0; i < rampUpDuration; i++ {
			load = append(load, profile.BaseLoad+(profile.MaxLoad-profile.BaseLoad)*i/rampUpDuration)
		}
		for i := rampUpDuration; i < duration; i++ {
			load = append(load, profile.MaxLoad)
		}
	case "spike":
		spikeDuration := duration / 4
		for i := 0; i < spikeDuration; i++ {
			load = append(load, profile.BaseLoad+(profile.MaxLoad-profile.BaseLoad)*i/spikeDuration)
		}
		for i := spikeDuration; i < 3*spikeDuration; i++ {
			load = append(load, profile.MaxLoad)
		}
		for i := 3 * spikeDuration; i < duration; i++ {
			load = append(load, profile.BaseLoad+(profile.MaxLoad-profile.BaseLoad)*(duration-i)/spikeDuration)
		}
	case "peak":
		peakDuration := duration / 3
		for i := 0; i < peakDuration; i++ {
			load = append(load, profile.BaseLoad+(profile.MaxLoad-profile.BaseLoad)*i/peakDuration)
		}
		for i := peakDuration; i < 2*peakDuration; i++ {
			load = append(load, profile.MaxLoad)
		}
		for i := 2 * peakDuration; i < duration; i++ {
			load = append(load, profile.MaxLoad-(profile.MaxLoad-profile.BaseLoad)*(i-2*peakDuration)/peakDuration)
		}
	}

	return load
}
