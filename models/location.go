package models

import (
	"math"
	"math/rand"
)

type (
	Location struct {
		Latitude  float64
		Longitude float64
	}
)

const (
	North = iota
	South
	East
	West
)

const (
	LATITUDE_PER_METRE = 1 / (111.32 * 1000)
)

func LONGITUDE_PER_METRE(l float64) float64 {
	return 1 / (40075 * 1000 * RadToDeg(math.Cos(DegToRad(l))) / 360)
}

func DegToRad(deg float64) float64 {
	return deg * 2 * (math.Phi / 360)
}

func RadToDeg(rad float64) float64 {
	return rad * (360 / 2 * math.Phi)
}

func ShiftLoc(l Location, dist float64, direction int) Location {
	switch direction {
	case North:
		return Location{
			Longitude: l.Longitude,
			Latitude:  l.Latitude + LATITUDE_PER_METRE*dist,
		}
	case South:
		return Location{
			Longitude: l.Longitude,
			Latitude:  l.Latitude - LATITUDE_PER_METRE*dist,
		}
	case East:
		newLong := l.Longitude + LONGITUDE_PER_METRE(l.Latitude)*dist
		return Location{
			Longitude: newLong,
			Latitude:  l.Latitude,
		}
	case West:
		newLong := l.Longitude - LONGITUDE_PER_METRE(l.Latitude)*dist
		return Location{
			Longitude: newLong,
			Latitude:  l.Latitude,
		}
	}
	return Location{}
}

func GenerateRandomLocation(l Location, dist float64) Location {
	if dist < 0 {
		return Location{}
	}

	dist = dist * rand.Float64()

	directions := []int{North, South, East, West}
	randDirection := directions[rand.Intn(len(directions))]
	location := ShiftLoc(l, dist, randDirection)

	return location
}
