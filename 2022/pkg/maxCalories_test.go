package maxCalories

import "testing"

func TestMaxCalories(t *testing.T) {
	calorieList := `1000
2000
3000

4000

5000
6000

7000
8000
9000

10000`
	expectedMax := 24_000
	maxCalories := FindMaxCalories(calorieList)

	if maxCalories != expectedMax {
		t.Errorf("expected %d but got %d", expectedMax, maxCalories)
	}
}
