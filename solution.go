package square

import (
	"math"
)


type figureSides int16

const SidesTriangle figureSides = 3
const SidesSquare figureSides = 4
const SidesCircle figureSides = 0


func CalcSquare(sideLen float64, sidesNum figureSides) float64 {
	if (sidesNum == SidesTriangle) {
		return ((math.Sqrt(3.) / 4.) * sideLen * sideLen)
	} else if (sidesNum == SidesSquare) {
		return (sideLen * sideLen)
	} else if (sidesNum == SidesCircle) {
		return (math.Pi * sideLen * sideLen);
	} else {
		return 0
	}
}
