package shapes

import (
	"math"
	"testing"
)

func TestArea(t *testing.T) {
	t.Run("Test Rectangle area", func(t *testing.T) {
		r := Rectangle{10, 10}
		expectedArea := float64(100)
		receivedArea := r.Area()

		if expectedArea != receivedArea {
			t.Errorf("Expected area is %f but receivedArea is %f", expectedArea, receivedArea)
		}
	})

	t.Run("Test Circle area", func(t *testing.T) {
		c := Circle{10}
		expectedArea := float64(math.Pi * (10 * 10))
		receivedArea := c.Area()

		if expectedArea != receivedArea {
			t.Errorf("Expected area is %f but receivedArea is %f", expectedArea, receivedArea)
		}
	})

}
