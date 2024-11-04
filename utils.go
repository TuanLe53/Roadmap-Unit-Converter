package main

import (
	"errors"
)

func convertLength(source_value float64, source_unit, target_unit string) (float64, error) {
	unitsInMeters := map[string]float64{
		"km":   1000,
		"hm":   100,
		"dam":  10,
		"m":    1,
		"dm":   0.1,
		"cm":   0.01,
		"mm":   0.001,
		"inch": 0.0254,
		"foot": 0.3048,
		"yard": 0.9144,
		"mile": 1609.34,
	}

	// Check if the units are supported
	fromFactor, ok1 := unitsInMeters[source_unit]
	toFactor, ok2 := unitsInMeters[target_unit]
	if !ok1 || !ok2 {
		return 0, errors.New("unsupported unit")
	}

	// Convert value to meters
	valueInMeters := source_value * fromFactor

	// Convert from meters to the target unit
	convertedValue := valueInMeters / toFactor

	return convertedValue, nil
}
