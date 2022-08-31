package scales

import (
	"errors"
)

type gender struct {
	boy  bool
	girl bool
}

func (g *gender) Boy() {
	g.boy = true
}

func (g *gender) Girl() {
	g.girl = true
}

func (g *gender) validate() error {
	if g.boy && g.girl {
		return errors.New("Both gender set up")
	}

	if !g.boy && !g.girl {
		return errors.New("No gender specified")
	}

	return nil
}
