package graphql

import (
	"fmt"

	"github.com/buoyantio/emojivoto/db"
	graphql "github.com/graph-gophers/graphql-go"
)

type (
	Resolver struct {
		db *db.DBClient
	}

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

func NewGraphQLServer(db *db.DBClient) *graphql.Schema {
	return graphql.MustParseSchema(Schema, &Resolver{db})
}
