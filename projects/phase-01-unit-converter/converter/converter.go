package converter

import (
	"fmt"
	"strings"
)

var supportedUnits = map[string]map[string]struct{}{
	"temperature": {"c": {}, "f": {}, "k": {}},
	"distance":    {"m": {}, "km": {}, "mi": {}},
	"weight":      {"g": {}, "kg": {}, "lb": {}},
}

func Convert(kind, from, to string, value float64) (float64, error) {
	kind = normalize(kind)
	from = normalize(from)
	to = normalize(to)

	units, ok := supportedUnits[kind]
	if !ok {
		return 0, fmt.Errorf("unsupported kind %q", kind)
	}
	if _, ok := units[from]; !ok {
		return 0, fmt.Errorf("unsupported source unit %q for kind %q", from, kind)
	}
	if _, ok := units[to]; !ok {
		return 0, fmt.Errorf("unsupported target unit %q for kind %q", to, kind)
	}
	if from == to {
		return value, nil
	}

	switch kind {
	case "temperature":
		return convertTemperature(from, to, value), nil
	case "distance":
		return convertDistance(from, to, value), nil
	case "weight":
		return convertWeight(from, to, value), nil
	default:
		return 0, fmt.Errorf("unsupported kind %q", kind)
	}
}

func normalize(s string) string {
	return strings.ToLower(strings.TrimSpace(s))
}

func convertTemperature(from, to string, value float64) float64 {
	celsius := value
	switch from {
	case "f":
		celsius = (value - 32) * 5 / 9
	case "k":
		celsius = value - 273.15
	}

	switch to {
	case "f":
		return celsius*9/5 + 32
	case "k":
		return celsius + 273.15
	default:
		return celsius
	}
}

func convertDistance(from, to string, value float64) float64 {
	meters := value
	switch from {
	case "km":
		meters = value * 1000
	case "mi":
		meters = value * 1609.344
	}

	switch to {
	case "km":
		return meters / 1000
	case "mi":
		return meters / 1609.344
	default:
		return meters
	}
}

func convertWeight(from, to string, value float64) float64 {
	kilograms := value
	switch from {
	case "g":
		kilograms = value / 1000
	case "lb":
		kilograms = value * 0.45359237
	}

	switch to {
	case "g":
		return kilograms * 1000
	case "lb":
		return kilograms / 0.45359237
	default:
		return kilograms
	}
}
