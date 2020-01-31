package sqrt

func almostEqual(v1, v2 float64) bool {
	return Abs(v1-v2) <= 0.001
}
