package readers

import (
	"bufio"
	"errors"
	"io"
	"os"

	"github.com/wispwisp/mmil/scales"
)

func read(in io.Reader, s *scales.Scales) error {
	// +1 for '\n', +1 for gender
	readLimit := scales.ANSWERSLIM + 2

	b1 := make([]byte, readLimit)
	readed, err := in.Read(b1)
	if err != nil {
		return err
	}

	if readed != readLimit {
		return errors.New("Expected 377 ansers and gender at the end")
	}

	// Set answers
	for i := 0; i < scales.ANSWERSLIM; i++ {
		switch b1[i] {
		case '0':
			s.RecieveAnswer(i, 0)
		case '1':
			s.RecieveAnswer(i, 1)
		default:
			return errors.New("Unexpeted symbol. Use (1-yes / 0-no) and (g-girl/b-boy) at the end")
		}
	}

	// Set gender
	switch b1[scales.ANSWERSLIM] {
	case 'g':
		s.SetGender().Girl()
	case 'b':
		s.SetGender().Boy()
	default:
		return errors.New("Unexpeted symbol. Use (1-yes / 0-no) and (g-girl/b-boy) at the end")
	}

	return nil
}

func ReadFromInput(s *scales.Scales) error {
	in := bufio.NewReader(os.Stdin)
	return read(in, s)
}

func ReadFromFile(fileName string, s *scales.Scales) error {
	f, err := os.Open(fileName)
	if err != nil {
		return err
	}

	in := bufio.NewReader(f)
	return read(in, s)
}
