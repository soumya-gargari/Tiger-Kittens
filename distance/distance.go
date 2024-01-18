package distance

import (
	"fmt"

	geo "github.com/kellydunn/golang-geo"
)

func CalculateDistance(lat1, lon1, lat2, lon2 float64) float64 {
	// Coordinates for Point A (latitude and longitude in decimal degrees)

	// Coordinates for Point B (latitude and longitude in decimal degrees)

	// Create GeoPoint for each location
	pointA := geo.NewPoint(lat1, lon1)
	pointB := geo.NewPoint(lat2, lon2)

	// Calculate the distance using Vincenty formula
	distance := pointA.GreatCircleDistance(pointB)

	// Print the result in kilometers
	fmt.Printf("Distance between the two points: %.2f kilometers\n", distance)
	return distance
}
