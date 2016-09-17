package calc

func round(x float64) float64 {
	if x > 0.5 {
		return float64(int(x + .5))
	}

	return float64(int(x - .5))
}
