package calc

func sign(x float64) float64 {
	if x > 0.0 {
		return 1.0
	} else if x < 0.0 {
		return -1.0
	}

	return 0.0
}
