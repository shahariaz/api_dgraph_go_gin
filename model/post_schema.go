package model

const PostSchema = `
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
`
