package main

import (
	"fmt"

	"github.com/wispwisp/mmil/readers"
	"github.com/wispwisp/mmil/scales"
)

func main() {
	s := scales.Scales{}

	// TODO: arg to specify kind of input
	if false {
		err := readers.ReadFromInput(&s)
		if err != nil {
			fmt.Println("Error reading answers from stdin:", err)
			return
		}
	} else if true {
		fileName := "./input.txt"
		err := readers.ReadFromFile(fileName, &s)
		if err != nil {
			fmt.Println("Error reading answers from file", fileName, ", error:", err)
			return
		}
	}

	res, err := s.Evaluate()
	if err != nil {
		fmt.Println("Error evaluating result:", err)
	} else {
		res.PrintScales()
	}
}
