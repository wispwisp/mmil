package scales

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type Result struct {
	WelshIndex int
	RawScales  [SCALESLIM]int
	TScales    [SCALESLIM]float64
}

func (r *Result) write(out io.Writer) {
	// TODO: gender in result
	fmt.Fprintln(out, "Welsh index:", r.WelshIndex)
	scaleNames := []string{"L", "F", "K", "1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}
	fmt.Fprintf(out, "Raw data | \t T-scales:\n")
	for i := 0; i < SCALESLIM; i++ {
		fmt.Fprintf(out, "%s= %v  \t | \t %s= %.0f\n", scaleNames[i], r.RawScales[i], scaleNames[i], r.TScales[i])
	}
}

func (r *Result) PrintScales() {
	r.write(bufio.NewWriter(os.Stdout))
}

func (r *Result) WriteToFile(fileName string) {
	f, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Failed to open file", fileName)
		return
	}
	defer f.Close()

	out := bufio.NewWriter(f)
	defer out.Flush()

	r.write(out)
}
