//go:build !seq
// +build !seq

package table

import (
	"sort"

	"github.com/sinux-l5d/INFO002-TP1/internal/config"
)

func NewTable(config config.Config, largeur uint64, hauteur uint64, random bool) (table, error) {
	T := make([][]uint64, hauteur)

	// Initial code was sequential, I rewrote it to be concurrent.
	// with size=2, abc=26, width=1000 and height=2000, I reduced the time from ~19s to ~5s
	type result struct {
		index uint64
		err   error
	}

	results := make(chan result)

	for i := range T {
		go func(idx uint64) {
			// INIT
			T[idx] = make([]uint64, 2)
			if random {
				T[idx][0] = index_aleatoire(&config)
			} else {
				T[idx][0] = uint64(idx)
			}

			// FILL
			var err error
			T[idx][1], err = nouvelle_chaine(&config, T[idx][0], largeur)
			results <- result{index: idx, err: err}
		}(uint64(i))
	}

	var errors []error

	for range T {
		res := <-results
		if res.err != nil {
			errors = append(errors, res.err)
		}
	}

	if len(errors) > 0 {
		return table{}, errors[0]
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
