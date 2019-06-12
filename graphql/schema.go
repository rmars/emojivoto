package graphql

import (
	"context"
	"fmt"

	"github.com/buoyantio/emojivoto/db"
	pb "github.com/buoyantio/emojivoto/emojivoto-web/gen/proto"
	graphql "github.com/graph-gophers/graphql-go"
)

type (
	Resolver struct {
		db                 *db.DBClient
		emojiServiceClient pb.EmojiServiceClient
	}

	helloResolver struct {
		Resolver
		string
	}

	userResolver struct {
		Resolver
		u *db.User
	}

	emojiResolver struct {
		Resolver
		emoji *pb.Emoji
	}
)

const Schema = `
schema {
	query: Query
}

type Query {
	hello(name: String!): String!
	users: [User]!
	emojis: [Emoji]!
	emoji(shortcode: String!): Emoji
}

type User {
	name: String!
	favEmoji: Emoji!
}

type Emoji {
	shortcode: String!
	unicode: String!
}
`

func (r *Resolver) Hello(args *struct{ Name string }) string {
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

func (r *userResolver) FavEmoji(ctx context.Context) (*emojiResolver, error) {
	emojiRsp, err := r.emojiServiceClient.FindByShortcode(ctx, &pb.FindByShortcodeRequest{
		Shortcode: r.u.FavEmoji,
	})

	return &emojiResolver{r.Resolver, emojiRsp.Emoji}, err
}

func (r *Resolver) Emojis(ctx context.Context) ([]*emojiResolver, error) {
	emojis := make([]*emojiResolver, 0)
	emojiRsp, err := r.emojiServiceClient.ListAll(ctx, &pb.ListAllEmojiRequest{})

	for _, emojiRsp := range emojiRsp.GetList() {
		emojis = append(emojis, &emojiResolver{*r, emojiRsp})
	}

	return emojis, err
}

func (r *Resolver) Emoji(ctx context.Context, args *struct{ Shortcode string }) (*emojiResolver, error) {
	emojiRsp, err := r.emojiServiceClient.FindByShortcode(ctx, &pb.FindByShortcodeRequest{
		Shortcode: args.Shortcode,
	})

	return &emojiResolver{*r, emojiRsp.Emoji}, err
}

func (r *emojiResolver) Unicode() string {
	return r.emoji.GetUnicode()
}

func (r *emojiResolver) Shortcode() string {
	return r.emoji.GetShortcode()
}

func NewGraphQLServer(db *db.DBClient, emojiServiceClient pb.EmojiServiceClient) *graphql.Schema {
	return graphql.MustParseSchema(Schema, &Resolver{db, emojiServiceClient})
}
