package external

import (
	"math"
	"testing"
)

func TestHaversine(t *testing.T) {
	withinPrecision(haversine(-50, 50, 0, 25), 6046e3, 10e3, t)
	withinPrecision(haversine(10, 10, -10, -10), 3137e3, 10e3, t)
	withinPrecision(haversine(90, 90, -90, -90), 20020e3, 10e3, t)
	withinPrecision(haversine(40, 73, -15, -35), 12590e3, 10e3, t)
	withinPrecision(haversine(-5, -5, 50, 50), 8063e3, 10e3, t)
	withinPrecision(haversine(-50, 50, -50, 50), 0, 10e3, t)
}

func withinPrecision(answer, target, precision float64, t *testing.T) {
	if math.Abs(target-answer) >= precision {
		t.Errorf(
			"provided answer not within acceptable precision. answer: %.2f, target: %.2f, precision: %.2f",
			answer,
			target,
			precision,
		)
	}
}
