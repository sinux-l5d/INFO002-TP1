//go:build seq
// +build seq

package table

import (
	"sort"

	"github.com/sinux-l5d/INFO002-TP1/internal/config"
)

func NewTable(config config.Config, largeur uint64, hauteur uint64, random bool) (table, error) {
	T := make([][]uint64, hauteur)

	for i := range T {
		// INIT
		T[i] = make([]uint64, 2)
		if random {
			T[i][0] = index_aleatoire(&config)
		} else {
			T[i][0] = uint64(i)
		}

		// FILL
		var err error
		T[i][1], err = nouvelle_chaine(&config, T[i][0], largeur)
		if err != nil {
			return table{}, err
		}
	}

	sort.Slice(T, func(i, j int) bool {
		return T[i][1] < T[j][1]
	})

	return table{
		Config:  config,
		Largeur: largeur,
		Hauteur: hauteur,
		Data:    T,
		Random:  random,
	}, nil
}
