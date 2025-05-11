package main

import (
	"os"
	"time"

	clockface "github.com/igolt/go-with-tests/math"
)

func main() {
	clockface.SVGWriter(os.Stdout, time.Now())
}
