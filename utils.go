package yageocoding

import "math"

// deg2rad converts degrees to radians
func deg2rad(deg float64) float64 {
	return deg * (math.Pi/180)
}