/*
 * Copyright 2019 Thibault NORMAND
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package badger

import (
	"context"

	"go.zenithar.org/pkg/db"
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
			if err == badger.ErrKeyNotFound {
				return db.ErrNoResult
			}
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
