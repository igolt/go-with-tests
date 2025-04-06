package main

import (
	"os"
	"time"

	clockface "example.com/hello/math"
)

func main() {
	clockface.SVGWriter(os.Stdout, time.Now())
}
