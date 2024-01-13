package tests

import (
	"encoding/binary"
	"errors"

	"github.com/sinux-l5d/INFO002-TP1/internal/config"
)

type H2ITest struct {
	config *config.Config
	hash   []byte
	// numéro de la colonne
	c uint64
}

func NewH2ITest(cfg *config.Config, hash []byte, c uint64) (*H2ITest, error) {
	if len(hash) != 20 {
		return nil, errors.New("invalid hash length for sha1")
	}
	return &H2ITest{
		config: cfg,
		hash:   hash,
		c:      c,
	}, nil
}

func (t H2ITest) Run() (uint64, error) {
	return (binary.LittleEndian.Uint64(t.hash[:8]) + t.c) % uint64(t.config.N()), nil
}
