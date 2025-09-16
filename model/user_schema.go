package model

const UserSchema = `
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
`
