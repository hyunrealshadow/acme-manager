package graphql

import (
	"acme-manager/ent"
	"acme-manager/graphql/generated"
	"acme-manager/graphql/model"
	"github.com/99designs/gqlgen/graphql"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

// Resolver is the resolver root.
type Resolver struct{ client *ent.Client }

// newSchema creates a graphql executable schema.
func newSchema(client *ent.Client) graphql.ExecutableSchema {
	c := generated.Config{Resolvers: &Resolver{client: client}}
	c.Directives.Model = model.ModelDirective
	return generated.NewExecutableSchema(c)
}
