package resolver

import (
	"activitydash/api/graph/generated"
	"activitydash/api/graph/validation"

	"github.com/99designs/gqlgen/graphql"
	"github.com/go-playground/validator/v10"
)

// Resolver is the resolver root.
type Resolver struct {
	validate *validator.Validate
}

// NewSchema creates a graphql executable schema.
func NewSchema() graphql.ExecutableSchema {
	return generated.NewExecutableSchema(generated.Config{
		Resolvers: &Resolver{
			validate: validation.CustomValidator(),
		},
	})
}
