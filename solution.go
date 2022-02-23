package solution

import (
	"math"
)

func calcSquare(sideLen float64, sidesNum int32) float64 {
	var area float64
	if sidesNum == 0 { // Circle, in this case sideLen specifies the radius
		area = math.Pi * math.Pow(sideLen, 2)
	} else if sidesNum == 3 { // Triangle
		area = math.Sqrt(3) / 2 * math.Pow(sideLen, 2)
	} else if sidesNum == 4 {
		area = math.Pow(sideLen, 2)
	} else {
		area = 0
	}

	return (area)

}
