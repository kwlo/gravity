package physics

import (
	"go.uber.org/zap"
)

type DistanceEnum int

const (
	Normalized       DistanceEnum = 0
	Yards            DistanceEnum = 1
	Miles            DistanceEnum = 2
	Inches           DistanceEnum = 3
	NormalizedToYard float64      = 10000
	YardToMile       float64      = 0.000568181818181818
	YardToInch       float64      = 36
)

func ConversionMap() []map[string]interface{} {
	zap.S().Infof("Converting fro")
	conversion := map[string][int] {
		"fdsfS": 100,
	}
	return conversion
}

func ConvertDistance(
	source float64,
	sourceType DistanceEnum,
	targetType DistanceEnum) float64 {

	zap.S().Infof("Converting from %s to %s", sourceType, targetType)

	if sourceType == targetType {
		return source
	}

	return 5
}
