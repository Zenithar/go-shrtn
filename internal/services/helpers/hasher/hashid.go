package hasher

import (
	"time"

	"github.com/sony/sonyflake"
	"github.com/speps/go-hashids"

	"golang.org/x/xerrors"
)

type simpleHasher struct {
	hids  *hashids.HashID
	flake *sonyflake.Sonyflake
}

// SimpleHasher returns an Hasher instance based on hashids
func SimpleHasher() Hasher {
	// Set up new hasher
	hd := hashids.NewData()
	// Set the salt to the one provided
	hd.Salt = "123456789"
	hd.MinLength = 5

	h, err := hashids.NewWithData(hd)
	if err != nil {
		panic(err)
	}

	// Initialize sonyflake
	sf := sonyflake.NewSonyflake(sonyflake.Settings{})

	return &simpleHasher{
		flake: sf,
		hids:  h,
	}
}

// -----------------------------------------------------------------------------

func (s *simpleHasher) Generate(salt string, t time.Time) (string, error) {
	n, err := s.flake.NextID()
	if err != nil {
		return "", xerrors.Errorf("failed to generate a sequence number: %w", err)
	}

	// Encode the given time as Unix timestamp to create the ID.
	id, err := s.hids.EncodeInt64([]int64{int64(n)})
	if err != nil {
		return "", xerrors.Errorf("failed to encode hash: %w", err)
	}

	return id, nil
}
