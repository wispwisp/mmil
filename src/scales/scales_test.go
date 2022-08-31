package scales

import (
	"fmt"
	"testing"
)

func TestRecieveAnswerArgumentRanges(t *testing.T) {
	var tests = []struct {
		i, a    int
		success bool
	}{
		{0, 0, true},
		{0, 1, true},
		{0, 2, false},
		{1, 0, true},
		{378, 0, false},
		{2, -2, false},
		{-1, 0, false},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("index:%d, answer:%d", tt.i, tt.a)
		t.Run(testname, func(t *testing.T) {
			s := Scales{}
			err := s.RecieveAnswer(tt.i, tt.a)
			if err != nil && tt.success {
				t.Errorf("Expect error")
			} else if err == nil && !tt.success {
				t.Errorf("Expect success")
			}
		})
	}
}
