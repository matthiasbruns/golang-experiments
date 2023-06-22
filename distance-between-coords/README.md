# distance between coords

## About

This projects implements a function that finds the two closest and two farthest points in a list of 2D coordinates.
We use different ways to represent a 2D coordinate:
- `Point` struct
- `[]int` slice
- `[2]int` array

to see, if there is a difference in performance and allocation.
The assumption is, that the `Point` struct and the `[2]int` array are the fastest and most efficient ways to represent a 2D coordinate.
This project should clarify, if pre-optimisation is necessary or if the compiler is smart enough to optimize the code.
Reading arrays is often harder for developers than reading structs, so the `Point` struct is the preferred way to represent a 2D coordinate.

## Golang Coding Challenge: Find the Two Closest and Two Farthest Points with Flexible Output

### Background

You have been given a list of 2D coordinates in a graph. Your task is to find both the two closest points and the two farthest points from this list using the `Point` struct. The result can be returned as either structs or arrays of coordinates.

### Task

First, create a `Point` struct representing a 2D coordinate:

```go
type Point struct {
	X int
	Y int
}
```

Next, write a Golang function with the following signature:

```go
func findClosestAndFarthestPoints(points []Point) (closestPoints, farthestPoints interface{})
```

**Input:**

- `points` is a slice of `n` `Point` instances where `2 <= n <= 10^3`, and `-10^4 <= points[i].X, points[i].Y <= 10^4`. Each point is unique.

**Output:**

- The function should return both `closestPoints` and `farthestPoints` as `interface{}` types. They can be either slices of `Point` structs or 2D arrays of `int` coordinates.
- `closestPoints` should contain the two points with the smallest Euclidean distance between them.
- `farthestPoints` should contain the two points with the largest Euclidean distance between them.
- If there is a tie for either the closest or farthest points, you may return any valid pair of closest or farthest points, respectively.

### Examples

```go
findClosestAndFarthestPoints([]Point{{0, 0}, {1, 1}, {3, 3}, {6, 6}})
/*
Returns:
closestPoints: [{0, 0}, {1, 1}] or [[0, 0], [1, 1]]
farthestPoints: [{0, 0}, {6, 6}] or [[0, 0], [6, 6]]
*/

findClosestAndFarthestPoints([]Point{{-5, 0}, {10, -2}, {6, 3}, {2, 4}})
/*
Returns:
closestPoints: [{6, 3}, {2, 4}] or [[6, 3], [2, 4]]
farthestPoints: [{-5, 0}, {10, -2}] or [[-5, 0], [10, -2]]
*/
```

### Note

You can use the standard `math` package for Golang to compute square roots or other mathematical operations. Make sure to handle the returned values with type assertions, as they are of type `interface{}`.