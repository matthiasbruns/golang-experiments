package main

import "testing"

func Test_findClosestAndFarthestPoints_1(t *testing.T) {
	points := [][2]int{{0, 0}, {1, 1}, {3, 3}, {6, 6}}
	closestPoints, farthestPoints := findClosestAndFarthestPoints(points)

	if closestPoints != [2][2]int{{0, 0}, {1, 1}} {
		t.Error("Expected closest points to be [[0, 0], [1, 1]]")
	}
	if farthestPoints != [2][2]int{{0, 0}, {6, 6}} {
		t.Error("Expected farthest points to be [[0, 0], [6, 6]]")
	}
}

func Test_findClosestAndFarthestPoints_2(t *testing.T) {
	points := [][2]int{{-5, 0}, {10, -2}, {6, 3}, {2, 4}}
	closestPoints, farthestPoints := findClosestAndFarthestPoints(points)

	if closestPoints != [2][2]int{{6, 3}, {2, 4}} {
		t.Error("Expected closest points to be [[6, 3], [2, 4]]")
	}
	if farthestPoints != [2][2]int{{-5, 0}, {10, -2}} {
		t.Error("Expected farthest points to be [[-5, 0], [10, -2]]")
	}
}

func Benchmark_findClosestAndFarthestPoints(b *testing.B) {
	for i := 0; i < b.N; i++ {
		points := [][2]int{{-5, 0}, {10, -2}, {6, 3}, {2, 4}}
		findClosestAndFarthestPoints(points)
	}
}

func Test_calculateDistance_1(t *testing.T) {
	p1 := [2]int{0, 0}
	p2 := [2]int{1, 1}

	distance := calculateDistance(p1, p2)

	if distance != 1.4142135623730951 {
		t.Error("Expected distance to be 1.4142135623730951")
	}
}

func Test_calculateDistance_2(t *testing.T) {
	p1 := [2]int{-2, -3}
	p2 := [2]int{3, 4}

	distance := calculateDistance(p1, p2)

	if distance != 8.602325267042627 {
		t.Error("Expected distance to be 8.602325267042627")
	}
}

func Benchmark_calculateDistance(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p1 := [2]int{i - 1, i - 1}
		p2 := [2]int{2 * i, i}

		calculateDistance(p1, p2)
	}
}
