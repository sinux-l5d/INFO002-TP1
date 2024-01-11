package tests

import (
	"fmt"
	"math"

	"github.com/sinux-l5d/INFO002-TP1/internal/config"
)

type I2CTest struct {
	config *config.Config
	n      int
}

// Convert number to the combinaison
// (size=4, n=0, abc=26) -> "aaaa"
// (size=4, n=1, abc=26) -> "aaab"
// (size=4, n=26, abc=26) -> "aaba"
func NewI2CTest(cfg *config.Config, n int) (*I2CTest, error) {
	if n < 0 || n > cfg.N()-1 {
		return nil, BoundaryError{0, cfg.N() - 1}
	}
	return &I2CTest{
		config: cfg,
		n:      n,
	}, nil
}

// perform the conversion
// Math made with the help with ChatGPT, implementation on my own
func (t *I2CTest) Run() error {
	ik := func(n, k, s int) int {
		div := math.Pow(float64(s), float64(k))
		return int(math.Floor(float64(n)/div)) % 26
	}
	r := ""
	s := len(t.config.Alphabet())
	for k := t.config.Size - 1; k >= 0; k-- {
		r += string(t.config.Alphabet()[ik(t.n, k, s)])
	}
	fmt.Printf("i2c(%d)=%s\n", t.n, r)
	return nil
}
