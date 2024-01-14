package tests

import (
	"fmt"
	"math"

	"github.com/sinux-l5d/INFO002-TP1/internal/config"
)

type I2CTest struct {
	config *config.Config
	i      uint64
}

// Convert number to the combinaison
// (size=4, n=0, abc=26) -> "aaaa"
// (size=4, n=1, abc=26) -> "aaab"
// (size=4, n=26, abc=26) -> "aaba"
func NewI2CTest(cfg *config.Config, i uint64) (*I2CTest, error) {
	if i > cfg.N()-1 {
		return nil, BoundaryError{0, cfg.N() - 1}
	}
	return &I2CTest{
		config: cfg,
		i:      i,
	}, nil
}

// perform the conversion
// Math made with the help with ChatGPT, implementation on my own
func (t *I2CTest) Run() (string, error) {

	r := ""
	s := len(t.config.Alphabet())

	ik := func(i uint64, k int, s int) int {
		div := math.Pow(float64(s), float64(k))
		return int(math.Floor(float64(i)/div)) % s
	}

	for k := t.config.Size - 1; k >= 0; k-- {
		r += string(t.config.Alphabet()[ik(t.i, k, s)])

		if t.config.Verbose {
			fmt.Printf("i2c(%d,%d,%d)=%d=%s\n", t.i, k, s, ik(t.i, k, s), r[len(r)-1:])
		}
	}

	return r, nil
}
