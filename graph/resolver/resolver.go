package resolver

import (
	"activitydash/api/graph/generated"

	"github.com/99designs/gqlgen/graphql"
)

// Resolver is the resolver root.
type Resolver struct{}

// NewSchema creates a graphql executable schema.
func NewSchema() graphql.ExecutableSchema {
	return generated.NewExecutableSchema(generated.Config{
		Resolvers: &Resolver{},
	})
}
