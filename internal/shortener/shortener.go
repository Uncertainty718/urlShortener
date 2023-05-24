package shortener

import (
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// var (
// 	charset  = os.Getenv("CHARSET")
// 	urlLen   = os.Getenv("LENGTH")
// 	myDomain = os.Getenv("DOMAIN")
// )

type Shortener struct {
	OriginalURL string     `json:"originalurl"`
	ShortURL    string     `json:"shorturl"`
	randomizer  *rand.Rand `json:"-"`
	charset     string     `json:"-"`
	urlLen      string     `json:"-"`
	myDomain    string     `json:"-"`
}

func NewShortener() *Shortener {
	return &Shortener{
		randomizer: rand.New(rand.NewSource(time.Now().UnixNano())),
		charset:    os.Getenv("CHARSET"),
		urlLen:     os.Getenv("LENGTH"),
		myDomain:   os.Getenv("DOMAIN"),
	}
}

func (s *Shortener) Shorten() {
	intLen, err := strconv.Atoi(s.urlLen)
	if err != nil {
		log.Fatal("I'm down ", err)
	}
	b := make([]byte, intLen)
	for i := range b {
		b[i] = s.charset[s.randomizer.Intn(len(s.charset))]
	}
	s.ShortURL = s.myDomain + string(b)
}

func (s *Shortener) Reshorten(short string) {
	s.Shorten()
	if s.ShortURL == short {
		s.Reshorten(s.ShortURL)
	}
}
