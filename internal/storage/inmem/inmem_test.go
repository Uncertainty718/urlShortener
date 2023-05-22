package storage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	storage *Inmem = NewInmem()
)

func TestSaveData(t *testing.T) {
	tests := []struct {
		name  string
		og    string
		short string
		err   error
	}{
		{
			name:  "Positive SaveData",
			og:    "https://google.com",
			short: "search",
			err:   nil,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := storage.SaveData(tc.og, tc.short)
			assert.Equal(t, tc.err, err)
		})
	}
}
