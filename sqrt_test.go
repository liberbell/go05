package sqrt

import "testing"

func almostEqual(v1, v2 float64) bool {
	return Abs(v1-v2) <= 0.001
}

func TestSimple(t *testing.T) {
	val, err := Sqrt(2)

	if err != nil {
		f.Fatalf("error in calculation - %s\n", err)
	}

	if !almostEqual(val, 1.414214) {
		f.Fatalf("bad value - %s\n", val)
	}
}
