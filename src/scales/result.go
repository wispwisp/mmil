package scales

import (
	"fmt"
)

type Result struct {
	WelshIndex int
	RawScales  [SCALESLIM]int
	TScales    [SCALESLIM]float64
}

func (r *Result) PrintScales() {
	scaleNames := []string{"L", "F", "K", "1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}
	fmt.Println("Raw data | \t T-scales:")
	for i := 0; i < SCALESLIM; i++ {
		fmt.Printf("%s= %v  \t | \t %s= %.0f\n", scaleNames[i], r.RawScales[i], scaleNames[i], r.TScales[i])
	}
}
