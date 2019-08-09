package v1

import (
	"context"

	apiv1 "go.zenithar.org/shrtn/internal/services/pkg/link/v1"
	linkv1 "go.zenithar.org/shrtn/pkg/gen/go/shrtn/link/v1"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

// Resolver is the root resolver implementation
type resolver struct {
	links apiv1.Link
}

// Resolver returns a root resolver instance
func Resolver(links apiv1.Link) ResolverRoot {
	return &resolver{
		links: links,
	}
}

// -----------------------------------------------------------------------------

// Mutation returns the mutation resolver.
func (r *resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}

// Query returns the query resolver
func (r *resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *resolver }

func (r *mutationResolver) Shorten(ctx context.Context, input linkv1.CreateRequest) (*linkv1.Link, error) {
	res, err := r.links.Create(ctx, &input)
	if err != nil {
		return nil, err
	}
	return res.Link, nil
}

type queryResolver struct{ *resolver }

func (r *queryResolver) Resolve(ctx context.Context, code string) (*linkv1.Link, error) {
	res, err := r.links.Resolve(ctx, &linkv1.ResolveRequest{
		Hash: code,
	})
	if err != nil {
		return nil, err
	}
	return res.Link, nil
}
