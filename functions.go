package formula

import "math"

// The following functions are automatically exposed to any formula.

// round ...
func round(params ...float64) float64 {
	return math.Round(params[0])
}

// floor ...
func floor(params ...float64) float64 {
	return math.Floor(params[0])
}

// ceil ...
func ceil(params ...float64) float64 {
	return math.Ceil(params[0])
}

// min ...
func min(params ...float64) float64 {
	return math.Min(params[0], params[1])
}

// max ...
func max(params ...float64) float64 {
	return math.Max(params[0], params[1])
}

// abs ...
func abs(params ...float64) float64 {
	return math.Abs(params[0])
}

// sqrt ...
func sqrt(params ...float64) float64 {
	return math.Sqrt(params[0])
}

// log ...
func log(params ...float64) float64 {
	return math.Log(params[0])
}

// pow ...
func pow(params ...float64) float64 {
	return math.Pow(params[0], params[1])
}

// sin ...
func sin(params ...float64) float64 {
	return math.Sin(params[0])
}

// cos ...
func cos(params ...float64) float64 {
	return math.Cos(params[0])
}

// tan ...
func tan(params ...float64) float64 {
	return math.Tan(params[0])
}