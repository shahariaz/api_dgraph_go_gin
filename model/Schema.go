package model

import (
	"context"
	"github.com/dgraph-io/dgo/v210"
	"github.com/dgraph-io/dgo/v210/protos/api"
	"log"
)

// SetupSchema sets up the Dgraph schema for Users, Posts, and Comments
func SetupSchema(dgraphClient *dgo.Dgraph) error {
	schema := `
type User {
	User.username: string
	User.email: string
	User.full_name: string
	User.bio: string
	User.created_at: datetime
	User.followers_count: int
	User.following_count: int
	User.posts_count: int
	User.is_verified: bool
	User.location: string
	User.posts: [Post]
	User.following: [User]
}

type Post {
	Post.post_id: string
	Post.content: string
	Post.created_at: datetime
	Post.likes_count: int
	Post.shares_count: int
	Post.comments_count: int
	Post.hashtags: [string]
	Post.is_promoted: bool
	Post.author: User
	Post.comments: [Comment]
}

type Comment {
	Comment.comment_id: string
	Comment.content: string
	Comment.created_at: datetime
	Comment.likes_count: int
	Comment.author: User
	Comment.post: Post
}

# User predicates
User.username: string @index(exact) @upsert .
User.email: string @index(exact) @upsert .
User.full_name: string @index(fulltext) .
User.bio: string @index(fulltext) .
User.created_at: datetime @index(day) .
User.followers_count: int @index(int) .
User.following_count: int @index(int) .
User.posts_count: int @index(int) .
User.is_verified: bool .
User.location: string @index(term) .
User.posts: [uid] @reverse .
User.following: [uid] @reverse .

# Post predicates
Post.post_id: string @index(exact) @upsert .
Post.content: string @index(fulltext) .
Post.created_at: datetime @index(day) .
Post.likes_count: int @index(int) .
Post.shares_count: int .
Post.comments_count: int .
Post.hashtags: [string] @index(term) .
Post.is_promoted: bool .
Post.author: uid @reverse .
Post.comments: [uid] @reverse .

# Comment predicates
Comment.comment_id: string @index(exact) @upsert .
Comment.content: string @index(fulltext) .
Comment.created_at: datetime @index(day) .
Comment.likes_count: int .
Comment.author: uid @reverse .
Comment.post: uid @reverse .
`

	ctx := context.Background()
	if err := dgraphClient.Alter(ctx, &api.Operation{Schema: schema}); err != nil {
		log.Printf("Error setting up schema: %v", err)
		return err
	}

	log.Println("Dgraph schema setup completed successfully.")
	return nil
}
