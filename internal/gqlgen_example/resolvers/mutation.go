package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/wricardo/gqlgen_example/internal/gqlgen_example/generated"
)

func (r *mutationResolver) Healthcheck(ctx context.Context, input string) (string, error) {
	return "healhy", nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
