package physics

import (
	"fmt"
)

// DistanceEnum Enumeration for distance (Normalized, Yards, Miles, Inches)
type DistanceEnum int

// Constants
const (
	Normalized       DistanceEnum = 0
	Yards            DistanceEnum = 1
	Miles            DistanceEnum = 2
	Inches           DistanceEnum = 3
	NormalizedToYard float64      = 10000
	YardToMile       float64      = 0.000568181818181818
	YardToInch       float64      = 36
)

var conversionMap map[string]float64 = map[string]float64{
	"0_1": NormalizedToYard,
	"1_0": 1 / NormalizedToYard,
	"0_2": NormalizedToYard * YardToMile,
	"2_0": 1 / NormalizedToYard * 1 / YardToMile,
	"1_2": YardToMile,
	"2_1": 1 / YardToMile,
	"0_3": NormalizedToYard * YardToInch,
	"3_0": 1 / NormalizedToYard * 1 / YardToInch,
	"1_3": YardToInch,
	"3_1": 1 / YardToInch,
	"2_3": 1 / YardToMile * YardToInch,
	"3_2": 1 / YardToInch * YardToMile,
}

// GetDistanceRate returns the rate source distance to target distance
func getDistanceRate(
	sourceType DistanceEnum,
	targetType DistanceEnum,
) (float64, bool) {
	result, ok := conversionMap[fmt.Sprintf("%d_%d", sourceType, targetType)]

	if ok {
		return result, true
	}

	return 0, false
}

// ConvertDistance converts source distance to target distance
func ConvertDistance(
	source float64,
	sourceType DistanceEnum,
	targetType DistanceEnum,
) float64 {
	if sourceType == targetType {
		return source
	}

	rate, _ := getDistanceRate(sourceType, targetType)

	return source * rate
}
