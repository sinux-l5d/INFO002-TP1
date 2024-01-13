package tests

import (
	"github.com/sinux-l5d/INFO002-TP1/internal/config"
)

type I2ITest struct {
	config *config.Config
	i      uint64
	// numéro de la colonne
	c uint64
}

func NewI2ITest(cfg *config.Config, i uint64, c uint64) (*I2ITest, error) {
	return &I2ITest{
		config: cfg,
		i:      i,
		c:      c,
	}, nil
}

func (t *I2ITest) Run() (uint64, error) {
	i2c, err := NewI2CTest(t.config, t.i)
	if err != nil {
		return 0, err
	}

	clair, err := i2c.Run()
	if err != nil {
		return 0, err
	}

	ht, err := NewHashTest(t.config, "sha1", clair)
	if err != nil {
		return 0, err
	}

	hash, err := ht.Run()
	if err != nil {
		return 0, err
	}

	h2i, err := NewH2ITest(t.config, hash, t.c)
	if err != nil {
		return 0, err
	}

	newI, err := h2i.Run()
	if err != nil {
		return 0, err
	}

	return newI, nil
}
