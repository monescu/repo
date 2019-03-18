package model

import "math"

type Point2D struct {
	X float64
	Y float64
}

type Point interface {
	Distance(other Point)float64
	Equals(other Point)bool
}

func (point Point2D) Distance(other Point2D)float64{
	return math.Sqrt((other.X-point.X)*(other.X-point.X) +
		(other.Y-point.Y)*(other.Y-point.Y))
}

func (point Point2D) Equals(other Point2D)bool {
	return point.X == other.X && point.Y == other.Y
}