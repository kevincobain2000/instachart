package pkg

type RadarChart struct {
}

func NewRadarChart() *RadarChart {
	return &RadarChart{}
}

func (c *RadarChart) GetIndicators(values [][]float64) []float64 {
	if len(values) == 0 {
		return nil
	}

	// Initialize maxValues with the first set of values
	maxValues := make([]float64, len(values[0]))
	copy(maxValues, values[0])

	// Iterate over each set of values
	for _, set := range values {
		for i, value := range set {
			if value > maxValues[i] {
				maxValues[i] = value
			}
		}
	}
	return maxValues
}
