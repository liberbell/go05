package main

func Abs(val float64) float64 {
	if val < 0 {
		return -val
	}
	return val
}

func Sqrt(val float64) (float64, error) {
	if val < 0.0 {
		return 0.0, ErrNegSqrt
	}

	if val == 0.0 {
		return 0.0, nil
	}
}
