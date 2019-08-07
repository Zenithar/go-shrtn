package repositories

import (
	"context"

	"go.zenithar.org/shrtn/internal/models"
)

// Link represents link persistence contract.
type Link interface {
	Create(ctx context.Context, entity *models.Link) error
	Get(ctx context.Context, hash string) (*models.Link, error)
	Delete(ctx context.Context, hash string) error
}
