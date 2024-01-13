package table

import (
	"bytes"
	"encoding/hex"
	"errors"
	"fmt"
	"sync"
	"sync/atomic"

	"github.com/sinux-l5d/INFO002-TP1/internal/tests"
)

// Crack the hash USING THE INTERNAL CONFIGURATION.
// Won't use the global config.
func (tab table) Crack(hash string) (clair string, err error) {
	// if panic in h2i/i2i, nicely return error
	// In order to return the error, we have to name the return values (see func definition)
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("error cracking hash %s: %v", hash, r)
		}
	}()

	H, err := hex.DecodeString(hash)
	if err != nil {
		return "", err
	}

	var wg sync.WaitGroup
	resultCh := make(chan string, 1)
	stopCh := make(chan struct{})
	var stopChClosed int32 // Indicateur atomique pour éviter la fermeture multiple du channel stopCh

	for t := tab.Largeur - 1; t > 0; t-- {
		wg.Add(1)
		go func(t uint64) {
			defer wg.Done()

			idx := tab.h2i(H, t)

			for i := t + 1; i < tab.Largeur; i++ {
				idx = tab.i2i(idx, i)
			}

			select {
			case <-stopCh:
				// Une solution a déjà été trouvée, arrêter cette goroutine
				return
			default:
			}

			a, b, ok := recherche(tab, tab.Hauteur, idx)
			if !ok {
				return
			}

			for i := a; i <= b; i++ {
				clair, ok := tab.verifie_candidat(H, t, tab.Data[i][0])
				if ok {
					select {
					case resultCh <- clair:
						// Une solution a été trouvée, signaler aux autres goroutines de s'arrêter
						// On ferme le channel que si ce n'est pas déjà fait
						if atomic.CompareAndSwapInt32(&stopChClosed, 0, 1) {
							close(stopCh)
						}
					default:
					}
					return
				}
			}
		}(t)
	}

	go func() {
		wg.Wait()
		close(resultCh)
		// On ferme le channel que si ce n'est pas déjà fait
		if atomic.CompareAndSwapInt32(&stopChClosed, 0, 1) {
			close(stopCh)
		}
	}()

	// If all goroutines are done unsuccessfully, it returns an error because the channel stopCh is closed and resultCh is empty
	select {
	case clair := <-resultCh:
		return clair, nil
	case <-stopCh:
		return "", errors.New("not found")
	}
}

// vérifie si un candidat est correct
//   - hash : empreinte à inverser
//   - t : numéro de la colonne où a été trouvé le candidat
//   - candicat : indice candidat (de la colonne terr)
//
// return : le texte clair obtenu et un booléen indiquant si le candidat est correct
func (tab table) verifie_candidat(hash []byte, t uint64, candidat uint64) (string, bool) {
	for i := uint64(1); i < t; i++ {
		candidat = tab.i2i(candidat, i)
	}
	clair := tab.i2c(candidat)
	return clair, bytes.Equal(hash, tab.h(clair))
}

// recherche dichotomique dans la table les premières et dernières lignes dont
// la seconde colonne est égale à idx
//   - table : table arc-en-ciel
//   - hauteur : nombre de chaines dans la table
//   - idx : indice à rechercher dans la dernière (deuxième) colonne
//   - a et b : (résultats) numéros des premières et dernières lignes dont les
//     dernières colonnes sont égale à idx
func recherche(tab table, hauteur uint64, idx uint64) (uint64, uint64, bool) {
	a := uint64(0)
	b := hauteur - 1

	for a <= b {
		m := (a + b) / 2
		if tab.Data[m][1] == idx {
			a = m
			b = m
			for a > 0 && tab.Data[a-1][1] == idx {
				a--
			}
			for b < hauteur-1 && tab.Data[b+1][1] == idx {
				b++
			}
			return a, b, true
		} else if tab.Data[m][1] < idx {
			a = m + 1
		} else {
			if m == 0 {
				break
			}
			b = m - 1
		}
	}
	return 0, 0, false
}

func (t table) h2i(hash []byte, column uint64) uint64 {
	// WARNING: use of config.GlobalConfig for simplicity, learning purpose only
	test, err := tests.NewH2ITest(&t.Config, hash, column)
	if err != nil {
		panic(err)
	}

	i, err := test.Run()
	if err != nil {
		panic(err)
	}

	return i
}

func (t table) i2i(i uint64, column uint64) uint64 {
	// WARNING: use of config.GlobalConfig for simplicity, learning purpose only
	test, err := tests.NewI2ITest(&t.Config, i, column)
	if err != nil {
		panic(err)
	}

	i, err = test.Run()
	if err != nil {
		panic(err)
	}

	return i
}

func (t table) i2c(i uint64) string {
	// WARNING: use of config.GlobalConfig for simplicity, learning purpose only
	test, err := tests.NewI2CTest(&t.Config, i)
	if err != nil {
		panic(err)
	}

	clair, err := test.Run()
	if err != nil {
		panic(err)
	}

	return clair
}

func (t table) h(clair string) []byte {
	// WARNING: use of config.GlobalConfig for simplicity, learning purpose only
	test, err := tests.NewHashTest(&t.Config, "sha1", clair)
	if err != nil {
		panic(err)
	}

	hash, err := test.Run()
	if err != nil {
		panic(err)
	}

	return hash
}
