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

	userResolver struct {
		Resolver
		u *db.User
	}
)

const Schema = `
schema {
	query: Query
}

type Query {
	hello(name: String!): String!
	users: [User]!
}

type User {
	name: String!
	favEmoji: String!
}
`

func (r *Resolver) Hello(args helloArgs) string {
	return fmt.Sprintf("Hello %s!", args.Name)
}

func (r *Resolver) Users() ([]*userResolver, error) {
	users, err := r.db.GetUsers()

	usersRsp := make([]*userResolver, 0)
	for _, u := range users {
		usersRsp = append(usersRsp, &userResolver{*r, u})
	}

	return usersRsp, err
}

func (r *userResolver) Name() string {
	return r.u.Name
}

func (r *userResolver) FavEmoji() string {
	return r.u.FavEmoji
}

func NewGraphQLServer(db *db.DBClient) *graphql.Schema {
	return graphql.MustParseSchema(Schema, &Resolver{db})
}
