package model

const CommentSchema = `
type Comment {
  Comment.comment_id: string
  Comment.content: string
  Comment.created_at: datetime
  Comment.likes_count: int
  Comment.author: User
  Comment.post: Post
}

# Comment predicates
Comment.comment_id: string @index(exact) @upsert .
Comment.content: string @index(fulltext) .
Comment.created_at: datetime @index(day) .
Comment.likes_count: int .
Comment.author: uid @reverse .
Comment.post: uid @reverse .
`
