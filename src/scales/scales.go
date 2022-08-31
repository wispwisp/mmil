package scales

import (
	"errors"
	"fmt"
	"math"
)

const ANSWERSLIM = 377
const SCALESLIM = 13

type Scales struct {
	answers [ANSWERSLIM]int
	aGender gender
}

func (s *Scales) RecieveAnswer(index int, answer int) error {
	if index < 0 || index >= ANSWERSLIM {
		return errors.New("invalid index")
	}

	if answer != 0 && answer != 1 {
		return errors.New(fmt.Sprintln("wrong answer:", answer))
	}

	s.answers[index] = answer

	return nil
}

// SetGender().Boy()
// SetGender().Girl()
func (s *Scales) SetGender() *gender {
	return &s.aGender
}

func (s *Scales) Evaluate() (res Result, err error) {
	err = s.aGender.validate()
	if err != nil {
		return
	}

	// Gender corrections for scales
	displacement := 0
	if s.aGender.girl {
		displacement = 13
	}

	// Answers evaluation
	s.rawScalesCount(res.RawScales[:], s.answers[:], displacement)

	// Welsh Index:
	res.WelshIndex = res.RawScales[1] - res.RawScales[2]

	// MMIL specific - K scale correct some other scales
	//   *(pmas+3) = *(pmas+3) + m_round((*(pmas+2) * 0.5)) ;
	res.RawScales[3] = res.RawScales[3] + int(math.Round(float64(res.RawScales[2])*0.5))
	//   *(pmas+6) = *(pmas+6) + m_round((*(pmas+2) * 0.4)) ;
	res.RawScales[6] = res.RawScales[6] + int(math.Round(float64(res.RawScales[2])*0.4))
	//   *(pmas+9) = *(pmas+9) + *(pmas+2) ;
	res.RawScales[9] = res.RawScales[9] + res.RawScales[2]
	//   *(pmas+10) = *(pmas+10) + *(pmas+2) ;
	res.RawScales[10] = res.RawScales[10] + res.RawScales[2]
	//   *(pmas+11) = *(pmas+11) + m_round((*(pmas+2) * 0.2)) ;
	res.RawScales[11] = res.RawScales[11] + int(math.Round(float64(res.RawScales[2])*0.2))

	// raw scales values in to T_values
	formulaValM := []float64{
		/*man*/ 3.944, 5.756, 15.744, 12.044, 20.252, 18.068, 21.232, 21.606, 9.240, 27.384, 26.704, 18.620, 26.766,
		/*woman*/ 4.456, 6.768, 14.567, 14.456, 24.656, 20.332, 21.976, 31.976, 10.493, 31.840, 29.556, 18.952, 30.376}

	formulaValS := []float64{
		/*man*/ 2.236, 2.921, 3.881, 3.283, 4.144, 4.438, 4.177, 3.913, 2.771, 4.791, 2.463, 4.005, 7.037,
		/*woman*/ 2.434, 3.183, 4.143, 4.470, 5.033, 5.176, 4.119, 3.820, 3.295, 4.772, 5.254, 3.816, 7.748}

	for i := 0; i < SCALESLIM; i++ {
		num := 10 * (float64(res.RawScales[i]) - formulaValM[displacement+i])
		denum := formulaValS[displacement+i]
		res.TScales[i] = 50 + (num / denum)
	}

	return
}
