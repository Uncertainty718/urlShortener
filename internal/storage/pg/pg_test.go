package storage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var storage *Postgres = NewPostgres()

func TestSaveData(t *testing.T) {
	tests := []struct {
		name  string
		og    string
		prep  int
		short string
		err   error
	}{
		{
			name:  "Positive SaveData",
			og:    "https://google.com",
			short: "search",
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			res, err := storage.SaveData(tc.og, tc.short)
			assert.Equal(t, tc.err, err)
			assert.Equal(t, tc.short, res)
		})
	}
}
