package blogpost

import "time"

type Blogpost struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	PostContent string    `json:"post_content"`
	Author      Author    `json:"author"`
	Tags        []string  `json:"tags"`
	Comments    []Comment `json:"comments,omitempty"`
	PublishDate time.Time `json:"publish_date"`
}

type Author struct {
	FullName string `json:"full_name"`
	Email    string `json:"email"`
}

type Comment struct {
	Author      Author `json:"author"`
	CommentBody string `json:"comment_body"`
}
