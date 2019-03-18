package main

import "math"

type Point2D struct {
	x float64
	y float64
}

type Point interface {
	distance(other Point)float64
	equals(other Point)bool
}

func (point Point2D) distance(other Point2D)float64{
	return math.Sqrt((other.x-point.x)*(other.x-point.x) +
		(other.y-point.y)*(other.y-point.y))
}

func (point Point2D) equals(other Point2D)bool {
	return point.x == other.x && point.y == other.y
}