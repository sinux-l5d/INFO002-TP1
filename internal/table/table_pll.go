//go:build !seq
// +build !seq

package table

import (
	"fmt"
	"runtime"
	"sort"

	"github.com/sinux-l5d/INFO002-TP1/internal/config"
)

var (
	workers = runtime.NumCPU()
)

type result struct {
	first, last uint64
	err         error
}

// worker is a function that will be run as a goroutine
// it takes jobs from the jobs channel and sends results to the results channel
// it will stop when the jobs channel is closed
func worker(config config.Config, largeur uint64, random bool, jobs <-chan uint64, results chan<- result) {
	var r result
	for j := range jobs {
		if random {
			r.first = index_aleatoire(&config)
		} else {
			r.first = uint64(j)
		}
		r.last, r.err = nouvelle_chaine(&config, r.first, largeur)
		results <- r
	}
}

func NewTable(config config.Config, largeur uint64, hauteur uint64, random bool) (table, error) {
	T := make([][]uint64, hauteur)

	jobs := make(chan uint64, hauteur)
	results := make(chan result, hauteur)

	if config.Verbose {
		fmt.Printf("Using %d workers\n", workers)
	}

	// Setup workers
	for w := 0; w < workers; w++ {
		go worker(config, largeur, random, jobs, results)
	}

	// Send jobs
	for i := range T {
		if config.Verbose {
			fmt.Printf("\rSending jobs %d%%", uint64(i*100)/hauteur)
		}
		// ensure all T[i] are initialized (although doing so in the collector should work)
		T[i] = make([]uint64, 2)
		jobs <- uint64(i) // send the index to the worker
	}

	if config.Verbose {
		fmt.Printf("\rSending jobs 100%%\n")
	}

	// Close jobs channel
	// It will stop the workers when they are done with their current job
	close(jobs)

	// Collect results
	for i := range T {
		if config.Verbose {
			fmt.Printf("\rProcessing %d%%", uint64(i*100)/hauteur)
		}

		res := <-results
		if res.err != nil {
			return table{}, res.err
		}

		// We don't need the index the worker have been working on, we'll sort anyway
		T[i][0] = res.first
		T[i][1] = res.last
	}

	if config.Verbose {
		fmt.Printf("\rProcessing 100%%\n")
	}

	// Sort based on the last element of each line
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
