package tests

import (
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"

	"github.com/sinux-l5d/INFO002-TP1/internal/config"
)

type H2ITest struct {
	config *config.Config
	hash   string
	// numéro de la colonne
	c int
}

func NewH2ITest(cfg *config.Config, hash string, c int) (*H2ITest, error) {
	if hash == "" {
		return nil, errors.New("empty string")
	}
	return &H2ITest{
		config: cfg,
		hash:   hash,
		c:      c,
	}, nil
}

func (t *H2ITest) Run() (uint64, error) {
	H, err := hex.DecodeString(t.hash)
	if err != nil {
		return 0, fmt.Errorf("invalid hash: %w", err)
	}

	return (binary.LittleEndian.Uint64(H[:8]) + uint64(t.c)) % uint64(t.config.N()), nil
}
