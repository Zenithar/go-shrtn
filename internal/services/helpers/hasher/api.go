package hasher

import "time"

// Hasher defines shortner hashing contract
type Hasher interface {
	Generate(salt string, t time.Time) (string, error)
}
