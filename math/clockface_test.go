package clockface

import (
	"fmt"
	"math"
	"testing"
	"time"
)

func TestSecondsInRadians(t *testing.T) {
	cases := []struct {
		time  time.Time
		angle float64
	}{
		{simpleTime(0, 0, 30), math.Pi},
		{simpleTime(0, 0, 0), 0},
		{simpleTime(0, 0, 45), (math.Pi / 2) * 3},
		{simpleTime(0, 0, 7), (math.Pi / 30) * 7},
	}

	for _, c := range cases {
		testName := fmt.Sprintf("convert %d second(s) to %g radians", c.time.Second(), c.angle)

		t.Run(testName, func(t *testing.T) {
			got := secondsInRadians(c.time)
			if got != c.angle {
				t.Fatalf("expected %v radians, but got %v", c.angle, got)
			}
		})
	}
}

func TestSecondsHandPoint(t *testing.T) {
	cases := []struct {
		time  time.Time
		point Point
	}{
		{simpleTime(0, 0, 30), Point{0, -1}},
		{simpleTime(0, 0, 45), Point{-1, 0}},
	}

	for _, c := range cases {
		testName := fmt.Sprintf("convert %d second(s) to position %v", c.time.Second(), c.point)

		t.Run(testName, func(t *testing.T) {
			got := secondHandPoint(c.time)
			if !roughlyEqualPoint(got, c.point) {
				t.Fatalf("expected %v Point, but got %v", c.point, got)
			}
		})
	}
}

func TestMinutesInRadians(t *testing.T) {
	cases := []struct {
		time  time.Time
		angle float64
	}{
		{simpleTime(0, 30, 0), math.Pi},
		{simpleTime(0, 0, 7), 7 * (math.Pi / (30 * 60))},
	}

	for _, c := range cases {
		testName := fmt.Sprintf(
			"convert %d minute(s) and %d second(s) to %g radians",
			c.time.Minute(),
			c.time.Second(),
			c.angle,
		)
		t.Run(testName, func(t *testing.T) {
			got := minutesInRadians(c.time)

			if !roughlyEqualFloat64(got, c.angle) {
				t.Fatalf("expected %v radians, but got %v", c.angle, got)
			}
		})
	}
}

func TestMinuteHandPoint(t *testing.T) {
	cases := []struct {
		time  time.Time
		point Point
	}{
		{simpleTime(0, 30, 0), Point{0, -1}},
		{simpleTime(0, 45, 0), Point{-1, 0}},
	}

	for _, c := range cases {

		testName := fmt.Sprintf(
			"convert %d minute(s) and %d second(s) to position %v",
			c.time.Minute(),
			c.time.Second(),
			c.point,
		)
		t.Run(testName, func(t *testing.T) {
			got := minuteHandPoint(c.time)
			if !roughlyEqualPoint(got, c.point) {
				t.Fatalf("expected %v point, but got %v", c.point, got)
			}
		})
	}
}

func TestHoursInRadians(t *testing.T) {
	cases := []struct {
		time  time.Time
		angle float64
	}{
		{simpleTime(6, 0, 0), math.Pi},
		{simpleTime(0, 0, 0), 0},
		{simpleTime(21, 0, 0), math.Pi * 1.5},
		{simpleTime(0, 1, 30), math.Pi / ((6 * 60 * 60) / 90)},
	}

	for _, c := range cases {
		testName := fmt.Sprintf(
			"convert %d hour(s), %d minute(s) and %d second(s) to %g radians",
			c.time.Hour(),
			c.time.Minute(),
			c.time.Second(),
			c.angle,
		)
		t.Run(testName, func(t *testing.T) {
			got := hoursInRadians(c.time)

			if !roughlyEqualFloat64(got, c.angle) {
				t.Fatalf("expected %v radians, but got %v", c.angle, got)
			}
		})
	}
}

func TestHourHandPoint(t *testing.T) {
	cases := []struct {
		time  time.Time
		point Point
	}{
		{simpleTime(6, 0, 0), Point{0, -1}},
		{simpleTime(18, 0, 0), Point{0, -1}},
		{simpleTime(21, 0, 0), Point{-1, 0}},
	}

	for _, c := range cases {
		testName := fmt.Sprintf(
			"convert %d hour(s), %d minute(s) and %d second(s) to position %v",
			c.time.Hour(),
			c.time.Minute(),
			c.time.Second(),
			c.point,
		)
		t.Run(testName, func(t *testing.T) {
			got := hourHandPoint(c.time)

			if !roughlyEqualPoint(got, c.point) {
				t.Fatalf("expected %v point, but got %v", c.point, got)
			}
		})
	}
}

func roughlyEqualPoint(p1, p2 Point) bool {
	return roughlyEqualFloat64(p1.X, p2.X) && roughlyEqualFloat64(p1.Y, p2.Y)
}

const float64EqualityTreshold = 1e-7

func roughlyEqualFloat64(f1, f2 float64) bool {
	return math.Abs(f1-f2) < float64EqualityTreshold
}

func simpleTime(hours, minutes, seconds int) time.Time {
	return time.Date(312, time.October, 28, hours, minutes, seconds, 0, time.UTC)
}
