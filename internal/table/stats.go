package table

import (
	"math"

	"github.com/sinux-l5d/INFO002-TP1/internal/config"
)

func (tab table) Coverage() float64 {
	return Coverage(tab.Config, tab.Largeur, tab.Hauteur)
}

func Coverage(cfg config.Config, largeur uint64, hauteur uint64) float64 {
	m := float64(hauteur)
	v := 1.0
	N := float64(cfg.N())
	for i := uint64(0); i < largeur; i++ {
		v = v * (1 - m/N)
		m = N * (1 - math.Exp(-m/N))
	}
	return 100 * (1 - v)
}
