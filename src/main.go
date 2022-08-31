package main

import (
	"flag"
	"fmt"

	"github.com/wispwisp/mmil/readers"
	"github.com/wispwisp/mmil/scales"
)

type Args struct {
	ReadFromFile  *bool
	InputFileName *string
}

func registerArgs() (args Args) {
	args.ReadFromFile = flag.Bool("fromfile", true, "read input from file")
	args.InputFileName = flag.String("filename", "input.txt", "input file name")
	flag.Parse()
	return
}

func main() {
	args := registerArgs()
	s := scales.Scales{}

	// TODO: arg to specify kind of input
	if !*args.ReadFromFile {
		err := readers.ReadFromInput(&s)
		if err != nil {
			fmt.Println("Error reading answers from stdin:", err)
			return
		}
	} else {
		err := readers.ReadFromFile(*args.InputFileName, &s)
		if err != nil {
			fmt.Println("Error reading answers from file", *args.InputFileName, ", error:", err)
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
