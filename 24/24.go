package main

import (
	"fmt"

	"math"
)

type Point struct {
	x float64
	y float64
}

func ConstructPoint(x float64, y float64) Point {
	var point_to_return Point
	point_to_return.x = x
	point_to_return.y = y
	return point_to_return
}

func FindDistance(p1 *Point, p2 *Point) float64 {
	x_dif_square := (p1.x - p2.x)
	x_dif_square = x_dif_square * x_dif_square

	y_dif_squre := (p1.y - p2.y)
	y_dif_squre = y_dif_squre * y_dif_squre

	return math.Sqrt((x_dif_square + y_dif_squre))
}

func main() {
	point_a := ConstructPoint(0.0, 0.0)
	point_b := ConstructPoint(123.0, -5.0)

	fmt.Println(FindDistance(&point_a, &point_b))
}
