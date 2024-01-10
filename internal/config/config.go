package config

import (
	"fmt"
	"io"
	"math"
	"sort"
	"strings"
)

var alphabets = map[string]string{
	"26":  "abcdefghijklmnopqrstuvwxyz",
	"26A": "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
	"36":  "abcdefghijklmnopqrstuvwxyz0123456789",
	"40":  "abcdefghijklmnopqrstuvwxyz0123456789,;:$",
	"52":  "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz",
	"62":  "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz",
	"66":  "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz,;:$",
}

// meant for CLI help
func Alphabets() string {
	var s []string
	for k, v := range alphabets {
		s = append(s, fmt.Sprintf("\t    %s: %s\n", k, v))
	}
	sort.Strings(s)
	return strings.Join(s, "")
}

type Config struct {
	CustomAlphabet string `json:"alphabet"`
	Abc            string `json:"abc"`
	Size           int    `json:"size"`
	Verbose        bool   `json:"verbose"`
	Writer         io.Writer
}

func (c Config) Printf(format string, a ...interface{}) (n int, err error) {
	return c.Writer.Write([]byte(fmt.Sprintf(format, a...)))
}

func (c Config) N() int {
	return int(math.Pow(float64(len(c.Alphabet())), float64(c.Size)))
}

func (c Config) Alphabet() string {
	if c.Abc != "" {
		abc, ok := alphabets[c.Abc]
		if !ok {
			return alphabets["26"]
		}
		return abc
	}

	return c.CustomAlphabet
}

func (c Config) String() string {
	return fmt.Sprintf("alphabet: %s\nsize: %d\nN: %d\nverbose: %t\n", c.Alphabet(), c.Size, c.N(), c.Verbose)
}
