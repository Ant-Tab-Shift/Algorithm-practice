package main

import "fmt"

const inf = 109

type Point struct {
	x, y int
}

func isReflected(points [][]int) bool {
	pointSet := make(map[Point]struct{}, len(points))
	minX, maxX := inf, -inf
	for _, point := range points {
		pointSet[Point{point[0], point[1]}] = struct{}{}
		minX = min(minX, point[0])
		maxX = max(maxX, point[0])
	}

	doubledLineXCoord := minX + maxX
	for point := range pointSet {
		reflectedPoint := Point{doubledLineXCoord - point.x, point.y}
		if _, ok := pointSet[reflectedPoint]; !ok {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(isReflected([][]int{{1, 1}, {-1, 1}}))
	fmt.Println(!isReflected([][]int{{1, 1}, {-1, -1}}))
}
