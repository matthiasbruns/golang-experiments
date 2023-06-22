## Golang Coding Challenge: Find the Two Closest and Two Farthest Points

### Background

You have been given a list of 2D coordinates in a graph. Your task is to find both the two closest points and the two farthest points from this list.

### Task

Write a Golang function with the following signature:

```go
func findClosestAndFarthestPoints(points [][]int) (closestPoints, farthestPoints [][]int)
```

**Input:**

- `points` is a list of 2D integer coordinates represented as `[n][2]int` with dimensions `n x 2` where `2 <= n <= 10^3`, and `-10^4 <= points[i][0], points[i][1] <= 10^4`. Each coordinate pair is unique.

**Output:**

- The function should return two separate lists `closestPoints` and `farthestPoints`, both containing two coordinate pairs each.
- `closestPoints` should contain the two points with the smallest Euclidean distance between them.
- `farthestPoints` should contain the two points with the largest Euclidean distance between them.
- If there is a tie for either the closest or farthest points, you may return any valid pair of closest or farthest points, respectively.

### Examples

```go
findClosestAndFarthestPoints([][]int{{0, 0}, {1, 1}, {3, 3}, {6, 6}})
/*
Returns:
closestPoints: [[0, 0], [1, 1]]
farthestPoints: [[0, 0], [6, 6]]
*/

findClosestAndFarthestPoints([][]int{{-5, 0}, {10, -2}, {6, 3}, {2, 4}})
/*
Returns:
closestPoints: [[6, 3], [2, 4]] // Note that these can be in any order
farthestPoints: [[-5, 0], [10, -2]] // Note that these can be in any order
*/
```

### Note

You can use the standard `math` package for Golang to compute square roots or other mathematical operations.