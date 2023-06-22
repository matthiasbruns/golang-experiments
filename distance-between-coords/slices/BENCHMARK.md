# Benchmarks

_Implementation uses arrays to represent coordinates._

## Tests

### calculateDistance

_Calculate distance between two coordinates with arrays_

**Command**

```bash
go test -bench=. -benchmem
```

**Output:**

```
goos: darwin
goarch: arm64
pkg: github.com/matthiasbruns/golang-experiments/distance-between-coords/slices
Benchmark_calculateDistance
Benchmark_calculateDistance-10    	1000000000	         0.3166 ns/op	       0 B/op	       0 allocs/op
PASS
```

**Conclusion:**

- The benchmark shows that the calculation of the distance between two coordinates is very fast.
- The benchmark also shows that the memory allocation is zero.

### findClosestAndFarthestPoints

_Find closest and farthest points with arrays_

**Command**

```bash
go test -bench=. -benchmem
```

**Output:**

```
goos: darwin
goarch: arm64
pkg: github.com/matthiasbruns/golang-experiments/distance-between-coords/slices
Benchmark_findClosestAndFarthestPoints
Benchmark_findClosestAndFarthestPoints-10    	57215967	        20.64 ns/op	       0 B/op	       0 allocs/op
PASS
```

**Conclusion:**

- The benchmark shows that the calculation of the closest and farthest points is fast enough for a small number of
  points.
- The benchmark also shows that the memory allocation is zero.

## Summary

The benchmarks show that the calculation of the distance between two coordinates and the calculation of the closest and
farthest points is fast enough for a small number of points. The benchmarks also show that the memory allocation is
zero when using arrays as coordinates.