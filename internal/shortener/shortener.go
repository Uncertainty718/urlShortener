package shortener

import (
	"math/rand"
	"time"
)

const (
	charset = "abcdefghijklmnopqrstuvwxyz" +
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_"
	urlLen   = 10
	myDomain = "https://short/"
)

type Shortener struct {
	OriginalURL string
	ShortURL    string
	randomizer  *rand.Rand
}

func NewShortener() *Shortener {
	return &Shortener{
		randomizer: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (s *Shortener) Shorten(originalURL string) {
	s.OriginalURL = originalURL
	b := make([]byte, urlLen)
	for i := range b {
		b[i] = charset[s.randomizer.Intn(len(charset))]
	}
	s.ShortURL = myDomain + string(b)
}
