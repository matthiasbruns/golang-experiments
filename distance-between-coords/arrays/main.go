package main

import (
	"fmt"
	"math"
)

func main() {
	points := [][2]int{{0, 0}, {1, 1}, {3, 3}, {6, 6}}

	closestPoints, farthestPoints := findClosestAndFarthestPoints(points)

	fmt.Println("Closest points:", closestPoints)
	fmt.Println("Farthest points:", farthestPoints)
}

func findClosestAndFarthestPoints(points [][2]int) (closestPoints, farthestPoints [2][2]int) {
	// brute force approach
	// iterate over all points and calculate distance between them
	// keep track of the closest and farthest points
	closestDistance := math.MaxFloat64
	farthestDistance := 0.0

	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			distance := calculateDistance(points[i], points[j])
			if distance < closestDistance {
				closestDistance = distance
				closestPoints[0] = points[i]
				closestPoints[1] = points[j]
			}
			if distance > farthestDistance {
				farthestDistance = distance
				farthestPoints[0] = points[i]
				farthestPoints[1] = points[j]
			}
		}
	}

	return closestPoints, farthestPoints
}

func calculateDistance(p1 [2]int, p2 [2]int) float64 {
	// euclidean distance

	deltaX := p1[0] - p2[0]
	deltaY := p1[1] - p2[1]

	return math.Sqrt(float64(deltaX*deltaX + deltaY*deltaY))
}
