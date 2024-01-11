package tests

import (
	"crypto/sha1"
	"errors"
	"fmt"

	"github.com/sinux-l5d/INFO002-TP1/internal/config"
)

// impl Test
type HashTest struct {
	config *config.Config
	alg    string
	str    string
}

func NewHashTest(cfg *config.Config, alg string, str string) (*HashTest, error) {
	if str == "" {
		return nil, errors.New("empty string")
	}
	if alg == "" {
		alg = "sha1"
	}
	if alg != "sha1" {
		return nil, errors.New("unsupported algorithm")
	}
	return &HashTest{
		config: cfg,
		alg:    alg,
		str:    str,
	}, nil
}

func (t *HashTest) Run() error {
	if t.alg != "sha1" {
		return errors.New("unsupported algorithm")
	}
	h := sha1.New()
	h.Write([]byte(t.str))
	fmt.Printf("%X (%s)\n", h.Sum(nil), t.str)
	return nil
}
