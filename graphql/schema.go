package graphql

import (
	"fmt"

	graphql "github.com/graph-gophers/graphql-go"
)

type (
	Resolver struct{}

	helloResolver struct {
		Resolver
		string
	}

	helloArgs struct {
		Name string
	}
)

const Schema = `
schema {
	query: Query
}

type Query {
	hello(name: String!): String!
}
`

func (r *Resolver) Hello(args helloArgs) string {
	return fmt.Sprintf("Hello %s!", args.Name)
}

func NewGraphQLServer() *graphql.Schema {
	return graphql.MustParseSchema(Schema, &Resolver{})
}
