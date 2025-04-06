package clockface_test

import (
	"bytes"
	"encoding/xml"
	"testing"
	"time"

	clockface "example.com/hello/math"
)

func TestSecondHandAtMidnight(t *testing.T) {
	tm := time.Date(1337, time.January, 1, 0, 0, 0, 0, time.UTC)

	expected := clockface.Point{X: 150, Y: 150 - 90}
	got := clockface.SecondHand(tm)

	assertPointsAreEqual(t, got, expected)
}

func TestSecondHandAt30Seconds(t *testing.T) {
	tm := time.Date(1337, time.January, 1, 0, 0, 30, 0, time.UTC)

	expected := clockface.Point{X: 150, Y: 150 + 90}
	got := clockface.SecondHand(tm)

	assertPointsAreEqual(t, got, expected)
}

func TestSVGWriterSecondHand(t *testing.T) {
	cases := []struct {
		time time.Time
		line Line
	}{
		{simpleTime(0, 0, 0), Line{150, 150, 150, 60}},
		{simpleTime(0, 0, 15), Line{150, 150, 240, 150}},
		{simpleTime(0, 0, 30), Line{150, 150, 150, 240}},
		{simpleTime(0, 0, 45), Line{150, 150, 60, 150}},
	}

	for _, c := range cases {
		testName := c.time.String()
		t.Run(testName, func(t *testing.T) {
			b := bytes.Buffer{}
			clockface.SVGWriter(&b, c.time)

			var svg SVG
			if err := xml.Unmarshal(b.Bytes(), &svg); err != nil {
				t.Fatalf("expected SVGWriter to produce a valid XML")
			}

			if !containsLine(c.line, svg.Line) {
				t.Errorf("expected to find the second hand line %+v, in the SVG lines %+v", c.line, svg.Line)
			}
		})
	}
}

func assertPointsAreEqual(t testing.TB, got, expected clockface.Point) {
	t.Helper()

	if got != expected {
		t.Errorf("expected %v, but got %v", expected, got)
	}
}

func containsLine(l Line, ls []Line) bool {
	for _, line := range ls {
		if line == l {
			return true
		}
	}
	return false
}

func simpleTime(hours, minutes, seconds int) time.Time {
	return time.Date(312, time.October, 28, hours, minutes, seconds, 0, time.UTC)
}

type SVG struct {
	XMLName xml.Name `xml:"svg"`
	Text    string   `xml:",chardata"`
	Xmlns   string   `xml:"xmlns,attr"`
	Width   string   `xml:"width,attr"`
	Height  string   `xml:"height,attr"`
	ViewBox string   `xml:"viewBox,attr"`
	Version string   `xml:"version,attr"`
	Circle  Circle   `xml:"circle"`
	Line    []Line   `xml:"line"`
}

type Circle struct {
	Cx float64 `xml:"cx,attr"`
	Cy float64 `xml:"cy,attr"`
	R  float64 `xml:"r,attr"`
}

type Line struct {
	X1 float64 `xml:"x1,attr"`
	Y1 float64 `xml:"y1,attr"`
	X2 float64 `xml:"x2,attr"`
	Y2 float64 `xml:"y2,attr"`
}
