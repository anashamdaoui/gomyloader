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
	switch profile.Type {
	case "fixed":
		return generateFixedProfile(profile)
	case "ramp_up":
		return generateRampUpProfile(profile)
	case "spike":
		return generateSpikeProfile(profile)
	case "peak":
		return generatePeakProfile(profile)
	default:
		return nil
	}
}

func generateFixedProfile(profile LoadProfile) []int {
	var load []int
	for i := 0; i < int(profile.Duration.Seconds()); i++ {
		load = append(load, profile.MaxLoad)
	}
	return load
}

func generateRampUpProfile(profile LoadProfile) []int {
	var load []int
	duration := int(profile.Duration.Seconds())
	rampUpDuration := duration / 2

	for i := 0; i < rampUpDuration; i++ {
		load = append(load, profile.BaseLoad+(profile.MaxLoad-profile.BaseLoad)*i/rampUpDuration)
	}
	for i := rampUpDuration; i < duration; i++ {
		load = append(load, profile.MaxLoad)
	}
	return load
}

func generateSpikeProfile(profile LoadProfile) []int {
	var load []int
	duration := int(profile.Duration.Seconds())
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
	return load
}

func generatePeakProfile(profile LoadProfile) []int {
	var load []int
	duration := int(profile.Duration.Seconds())
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
	return load
}
