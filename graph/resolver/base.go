package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"activitydash/api/graph/generated"
	"activitydash/api/graph/model"
	"context"
	"fmt"
)

func (r *mutationResolver) Say(ctx context.Context, something string) (*model.SayPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Hello(ctx context.Context) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
