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

package v1

import (
	"context"
	"time"

	"go.zenithar.org/shrtn/internal/models"
	"go.zenithar.org/shrtn/internal/repositories"
	"go.zenithar.org/shrtn/internal/services/internal/hasher"
	linkv1 "go.zenithar.org/shrtn/pkg/gen/go/shrtn/link/v1"

	"golang.org/x/xerrors"
)

// Link defines link management service contract
type Link interface {
	Create(ctx context.Context, req *linkv1.CreateRequest) (*linkv1.CreateResponse, error)
	Resolve(ctx context.Context, req *linkv1.ResolveRequest) (*linkv1.ResolveResponse, error)
	Delete(ctx context.Context, req *linkv1.DeleteRequest) (*linkv1.DeleteResponse, error)
}

type service struct {
	h     hasher.Hasher
	links repositories.Link
}

// New returns a service instance
func New(links repositories.Link, h hasher.Hasher) Link {
	return &service{
		links: links,
		h:     h,
	}
}

// -----------------------------------------------------------------------------

// Create a shortened url entry.
func (s *service) Create(ctx context.Context, req *linkv1.CreateRequest) (*linkv1.CreateResponse, error) {
	// Validate request
	if err := req.Validate(); err != nil {
		return nil, xerrors.Errorf("link: unable to validate request: %w", err)
	}

	// Generate an identifier
	id, err := s.h.Generate(req.Url, time.Now())
	if err != nil {
		return nil, err
	}

	// Initialize an entity
	entity := &models.Link{
		Hash:        id,
		Url:         req.Url,
		Description: req.Description,
	}

	// Store in persistence
	if err := s.links.Create(ctx, entity); err != nil {
		return nil, err
	}

	// Return result
	return &linkv1.CreateResponse{
		Link: entity,
	}, nil
}

// Resolve a shortened url entry.
func (s *service) Resolve(ctx context.Context, req *linkv1.ResolveRequest) (*linkv1.ResolveResponse, error) {
	// Validate request
	if err := req.Validate(); err != nil {
		return nil, xerrors.Errorf("link: unable to validate request: %w", err)
	}

	// Check in persistence for existence
	entity, err := s.links.Get(ctx, req.Hash)
	if err != nil {
		return nil, xerrors.Errorf("link: unable to retrieve object from database: %w", err)
	}

	// Return result
	return &linkv1.ResolveResponse{
		Link: entity,
	}, nil
}

// Delete a shortened url entry.
func (s *service) Delete(ctx context.Context, req *linkv1.DeleteRequest) (*linkv1.DeleteResponse, error) {
	// Validate request
	if err := req.Validate(); err != nil {
		return nil, xerrors.Errorf("link: unable to validate request: %w", err)
	}

	// Delete from persistence
	if err := s.links.Delete(ctx, req.Hash); err != nil {
		return nil, xerrors.Errorf("link: unable to delete object from database: %w", err)
	}

	// Return result
	return &linkv1.DeleteResponse{}, nil
}
