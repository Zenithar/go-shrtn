package badger

import (
	"context"

	"go.zenithar.org/shrtn/internal/models"
	"go.zenithar.org/shrtn/internal/repositories"

	badger "github.com/dgraph-io/badger"
	"github.com/gogo/protobuf/proto"
	"golang.org/x/xerrors"
)

type badgerLinkRepository struct {
	db *badger.DB
}

// Link returns a link repository instance for Badger
func Link(db *badger.DB) repositories.Link {
	return &badgerLinkRepository{
		db: db,
	}
}

// -----------------------------------------------------------------------------

func (r *badgerLinkRepository) Create(ctx context.Context, entity *models.Link) error {
	// Validate entity
	if err := entity.Validate(); err != nil {
		return xerrors.Errorf("link: invalid entity: %w", err)
	}

	// RW Transaction
	return r.db.Update(func(txn *badger.Txn) error {
		// Serialize to protobuf
		payload, err := proto.Marshal(entity)
		if err != nil {
			return err
		}

		// Assign key
		return txn.Set([]byte(entity.Hash), payload)
	})
}

func (r *badgerLinkRepository) Get(ctx context.Context, hash string) (*models.Link, error) {
	var entity models.Link

	// RO Transaction
	if err := r.db.View(func(txn *badger.Txn) error {
		// Retrieve with key
		item, err := txn.Get([]byte(hash))
		if err != nil {
			return err
		}

		// Decode value
		return item.Value(func(val []byte) error {
			// Decode from protobuf
			return proto.Unmarshal(val, &entity)
		})
	}); err != nil {
		return nil, xerrors.Errorf("link: unable to retrieve from database: %w", err)
	}

	// return result
	return &entity, nil
}

func (r *badgerLinkRepository) Delete(ctx context.Context, hash string) error {
	// RW Transaction
	return r.db.Update(func(txn *badger.Txn) error {
		// Assign key
		return txn.Delete([]byte(hash))
	})
}
