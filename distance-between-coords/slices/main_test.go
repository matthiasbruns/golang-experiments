package main

import (
	"testing"
)

func Test_findClosestAndFarthestPoints_1(t *testing.T) {
	points := [][]int{{0, 0}, {1, 1}, {3, 3}, {6, 6}}
	closestPoints, farthestPoints := findClosestAndFarthestPoints(points)

	if closestPoints[0][0] != 0 || closestPoints[0][1] != 0 {
		t.Error("Expected closest point to be [0, 0]")
	}
	if closestPoints[1][0] != 1 || closestPoints[1][1] != 1 {
		t.Error("Expected closest point to be [1, 1]")
	}
	if farthestPoints[0][0] != 0 || farthestPoints[0][1] != 0 {
		t.Error("Expected farthest point to be [6, 6]")
	}
	if farthestPoints[1][0] != 6 || farthestPoints[1][1] != 6 {
		t.Error("Expected farthest point to be [3, 3]")
	}
}

func Test_findClosestAndFarthestPoints_2(t *testing.T) {
	points := [][]int{{-5, 0}, {10, -2}, {6, 3}, {2, 4}}
	closestPoints, farthestPoints := findClosestAndFarthestPoints(points)

	if closestPoints[0][0] != 6 || closestPoints[0][1] != 3 {
		t.Error("Expected closest point to be [6, 3]")
	}
	if closestPoints[1][0] != 2 || closestPoints[1][1] != 4 {
		t.Error("Expected closest point to be [2, 4]")
	}
	if farthestPoints[0][0] != -5 || farthestPoints[0][1] != 0 {
		t.Error("Expected farthest point to be [-5, 0]")
	}
	if farthestPoints[1][0] != 10 || farthestPoints[1][1] != -2 {
		t.Error("Expected farthest point to be [3, 3]")
	}
}

func Benchmark_findClosestAndFarthestPoints(b *testing.B) {
	for i := 0; i < b.N; i++ {
		points := [][]int{{-5, 0}, {10, -2}, {6, 3}, {2, 4}}
		findClosestAndFarthestPoints(points)
	}
}

func Test_calculateDistance_1(t *testing.T) {
	p1 := []int{0, 0}
	p2 := []int{1, 1}

	distance := calculateDistance(p1, p2)

	if distance != 1.4142135623730951 {
		t.Error("Expected distance to be 1.4142135623730951")
	}
}

func Test_calculateDistance_2(t *testing.T) {
	p1 := []int{-2, -3}
	p2 := []int{3, 4}

	distance := calculateDistance(p1, p2)

	if distance != 8.602325267042627 {
		t.Error("Expected distance to be 8.602325267042627")
	}
}

func Benchmark_calculateDistance(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p1 := []int{i - 1, i - 1}
		p2 := []int{2 * i, i}

		calculateDistance(p1, p2)
	}
}
