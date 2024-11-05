package main

import (
	"errors"
)

func convertLength(source_value float64, source_unit, target_unit string) (float64, error) {
	if source_unit == target_unit {
		return source_value, nil // No conversion needed if units are the same
	}

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

func convertWeight(source_value float64, source_unit, target_unit string) (float64, error) {
	if source_unit == target_unit {
		return source_value, nil // No conversion needed if units are the same
	}

	unitsInGrams := map[string]float64{
		"ton":   1_000_000, // metric ton (1 ton = 1,000,000 grams)
		"kg":    1000,
		"gram":  1,
		"mg":    0.001,
		"pound": 453.592,
		"ounce": 28.3495,
	}

	// Check if the units are supported
	fromFactor, ok1 := unitsInGrams[source_unit]
	toFactor, ok2 := unitsInGrams[target_unit]
	if !ok1 || !ok2 {
		return 0, errors.New("unsupported unit")
	}

	// Convert value to grams
	valueInGrams := source_value * fromFactor

	// Convert from grams to the target unit
	convertedValue := valueInGrams / toFactor

	return convertedValue, nil
}

func convertTemperature(source_value float64, source_unit, target_unit string) (float64, error) {
	if source_unit == target_unit {
		return source_value, nil // No conversion needed if units are the same
	}

	var result float64

	switch source_unit {
	case "Celsius":
		if target_unit == "Fahrenheit" {
			result = (source_value * 9 / 5) + 32
		} else if target_unit == "Kelvin" {
			result = source_value + 273.15
		} else {
			return 0, errors.New("unsupported target unit")
		}
	case "Fahrenheit":
		if target_unit == "Celsius" {
			result = (source_value - 32) * 5 / 9
		} else if target_unit == "Kelvin" {
			result = (source_value-32)*5/9 + 273.15
		} else {
			return 0, errors.New("unsupported target unit")
		}
	case "Kelvin":
		if target_unit == "Celsius" {
			result = source_value - 273.15
		} else if target_unit == "Fahrenheit" {
			result = (source_value-273.15)*9/5 + 32
		} else {
			return 0, errors.New("unsupported target unit")
		}
	default:
		return 0, errors.New("unsupported source unit")
	}

	return result, nil
}
