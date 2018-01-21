package main

import (
	"math"
	"fmt"
	"math/rand"
	"time"
)

const radius = 1

type Point struct {
	X, Y float64
}

func (p1 Point) Distance(p2 Point) (D float64) {
	deltaX := math.Abs(p1.X-p2.X)
	deltaY := math.Abs(p1.Y-p2.Y)
	return math.Sqrt(deltaX*deltaX + deltaY*deltaY)
}

func (p Point) IsInsideCircle() bool {
	if p.Distance(Point{0,0}) <= radius {
		return true
	}
	return false
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	trials := 0
	successes := 0
	var ratio float64

	for {
		trials++
		var p = Point{rand.Float64(),rand.Float64()}
		if p.IsInsideCircle() {
			successes++
		}
		time.Sleep(time.Millisecond*250)
		ratio = float64(successes)/float64(trials)
		fmt.Printf("Trials: %v Successes: %v Ratio: %v PI approx: %v\n", trials, successes, ratio, ratio*4)
	}
}
