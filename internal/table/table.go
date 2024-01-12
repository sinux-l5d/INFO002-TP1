package table

import (
	"encoding/gob"
	"fmt"
	"math/rand"
	"os"
	"sort"

	"github.com/sinux-l5d/INFO002-TP1/internal/config"
	"github.com/sinux-l5d/INFO002-TP1/internal/tests"
)

func nouvelle_chaine(config *config.Config, idx uint64, largeur uint64) (uint64, error) {
	id := idx

	for i := uint64(0); i < largeur-1; i++ {
		i2i, err := tests.NewI2ITest(config, id, i+1)
		if err != nil {
			return 0, err
		}

		id, err = i2i.Run()
		if err != nil {
			return 0, err
		}

		if config.Verbose {
			fmt.Printf("i2i(%d,%d)=%d | ", idx, i+1, id)
		}
	}
	if config.Verbose {
		fmt.Printf("i2i(%d,%d)=%d\n", idx, largeur, id)
	}

	return id, nil
}

func index_aleatoire(config *config.Config) uint64 {
	return rand.Uint64() % config.N() // bias ?
}

type table struct {
	Config  config.Config
	Hauteur uint64
	Largeur uint64
	Random  bool

	Data [][]uint64
}

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
			return table{}, nil
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

func (t table) Save(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("error creating file %s: %w", filename, err)
	}
	defer file.Close()

	enc := gob.NewEncoder(file)

	err = enc.Encode(t)
	if err != nil {
		return fmt.Errorf("error encoding file %s: %w", filename, err)
	}

	return nil
}

// Load a table with it's own configuration
func Load(filename string) (table, error) {
	file, err := os.Open(filename)
	if err != nil {
		return table{}, fmt.Errorf("error opening file %s: %w", filename, err)
	}
	defer file.Close()

	dec := gob.NewDecoder(file)

	var T table
	err = dec.Decode(&T)
	if err != nil {
		return table{}, fmt.Errorf("error decoding file %s: %w", filename, err)
	}

	return T, nil
}

func (t table) Print(n int) string {

	if n == 0 {
		n = len(t.Data)
	}

	r := ""
	for i := range t.Data {
		r += fmt.Sprintf("%06d : %d %d\n", i, t.Data[i][0], t.Data[i][1])
		if i == n {
			break
		}
	}
	return r
}
