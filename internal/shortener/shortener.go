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
	OriginalURL string     `json:"originalurl"`
	ShortURL    string     `json:"shorturl"`
	randomizer  *rand.Rand `json:"-"`
}

func NewShortener() *Shortener {
	return &Shortener{
		randomizer: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (s *Shortener) Shorten() {
	b := make([]byte, urlLen)
	for i := range b {
		b[i] = charset[s.randomizer.Intn(len(charset))]
	}
	s.ShortURL = myDomain + string(b)
}

func (s *Shortener) Reshorten(short string) {
	s.Shorten()
	if s.ShortURL == short {
		s.Reshorten(s.ShortURL)
	}
}
