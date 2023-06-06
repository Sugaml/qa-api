package entity

import "time"

type Comment struct {
	ID        string    // Unique identifier for the comment
	Author    string    // Name or identifier of the user who wrote the comment
	Content   string    // Content of the comment
	CreatedAt time.Time // Timestamp indicating when the comment was created
}
