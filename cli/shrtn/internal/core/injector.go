package core

import (
	bdo "github.com/dgraph-io/badger"
	"github.com/google/wire"

	"go.zenithar.org/shrtn/cli/shrtn/internal/config"
	"go.zenithar.org/shrtn/internal/repositories/pkg/badger"
	"go.zenithar.org/shrtn/internal/services/helpers/hasher"
	v1link "go.zenithar.org/shrtn/internal/services/pkg/link/v1"
)

// -----------------------------------------------------------------------------

// BadgerConfig initializes the database connection to badger KV Store.
func BadgerConfig(cfg *config.Configuration) (*bdo.DB, error) {
	return bdo.Open(bdo.DefaultOptions(cfg.DB.Path))
}

var repositorySet = wire.NewSet(
	BadgerConfig,
	badger.RepositorySet,
)

// -----------------------------------------------------------------------------

// Hasher declares an url hasher implementation
func Hasher(_ *config.Configuration) hasher.Hasher {
	return hasher.SimpleHasher()
}

// -----------------------------------------------------------------------------

// V1Services is the v1 services object injector.
var V1Services = wire.NewSet(
	repositorySet,
	Hasher,
	v1link.New,
)
