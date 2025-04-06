package clockface

import (
	"math"
	"time"
)

type Point struct {
	X float64
	Y float64
}

const secondHandLength = 90

const (
	clockCentreX float64 = 150
	clockCentreY float64 = 150
)

func SecondHand(t time.Time) Point {
	p := SecondHandPoint(t)
	return Point{
		p.X*secondHandLength + clockCentreX,
		-p.Y*secondHandLength + clockCentreY,
	}
}

func SecondsInRadians(t time.Time) float64 {
	return (math.Pi / (30 / float64(t.Second())))
}

func SecondHandPoint(t time.Time) Point {
	alpha := SecondsInRadians(t)
	return Point{math.Sin(alpha), math.Cos(alpha)}
}
