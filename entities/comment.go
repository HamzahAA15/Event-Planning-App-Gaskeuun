package entities

// dipisah folder pisah file juga
type Comment struct {
	Id      int    `json:"id" form:"id"`
	EventId int    `json:"event_id" form:"event_id"`
	UserId  int    `json:"user_id" form:"user_id"`
	Comment string `json:"comment" form:"comment"`
}

type CommentResponse struct {
	Id        int    `json:"id" form:"id"`
	EventId   int    `json:"event_id" form:"event_id"`
	User      User   `json:"user" form:"user"`
	Comment   string `json:"comment" form:"comment"`
	UpdatedAt string `json:"updated_at" form:"updated_at"`
}
