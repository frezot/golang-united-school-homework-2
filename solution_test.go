package square

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestCalcSquare(t *testing.T) {

	assert.Equal(t, CalcSquare(0, SidesTriangle), 0.0)
	assert.Equal(t, CalcSquare(1, SidesTriangle), 0.4330127018922193)
	assert.Equal(t, CalcSquare(21, SidesTriangle), 190.95860153446873)
	
	assert.Equal(t, CalcSquare(0, SidesSquare), 0.0)
	assert.Equal(t, CalcSquare(2, SidesSquare), 4.0)
	assert.Equal(t, CalcSquare(42, SidesSquare), 1764.0)
	assert.Equal(t, CalcSquare(65_000, SidesSquare), 4.225e+09)

	assert.Equal(t, CalcSquare(0, SidesCircle), 0.0)
	assert.Equal(t, CalcSquare(0, SidesCircle), 0.0)
	assert.Equal(t, CalcSquare(1, SidesCircle), math.Pi)
	assert.Equal(t, CalcSquare(36, SidesCircle), 4071.5040790523717)
	
	assert.Equal(t, CalcSquare(13, 02), 0.0)
	assert.Equal(t, CalcSquare(32, 42), 0.0)
	assert.Equal(t, CalcSquare(64, 06), 0.0)
}
