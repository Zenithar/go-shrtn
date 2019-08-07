package hasher

import (
	"time"

	"github.com/speps/go-hashids"
	"golang.org/x/xerrors"
)

type simpleHasher struct{}

// SimpleHasher returns an Hasher instance based on hashids
func SimpleHasher() Hasher {
	return &simpleHasher{}
}

// -----------------------------------------------------------------------------

func (s *simpleHasher) Generate(salt string, t time.Time) (string, error) {
	// Set up new hasher
	hd := hashids.NewData()
	// Set the salt to the one provided
	hd.Salt = salt

	// create new HashID
	h, err := hashids.NewWithData(hd)
	if err != nil {
		return "", xerrors.Errorf("failed to create new hash: %w", err)
	}

	// Encode the given time as Unix timestamp to create the ID.
	id, err := h.Encode([]int{int(t.Unix())})
	if err != nil {
		return "", xerrors.Errorf("failed to encode hash: %w", err)
	}

	return id, nil
}
