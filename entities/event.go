package entities

type Event struct {
	Id          int    `json:"id" form:"id"`
	UserID      int    `json:"userid" form:"userid"`
	CategoryId  int    `json:"categoryid" form:"categoryid"`
	Title       string `json:"title" form:"title"`
	Host        string `json:"host" form:"host"`
	Date        string `json:"date" form:"date"`
	Location    string `json:"location" form:"location"`
	Description string `json:"description" form:"description"`
	ImageUrl    string `json:"imageurl" form:"imageurl"`
}

type EventIdResponse struct {
	Event
	Participants User
	Comments     Comment
}

type Categories struct {
	Id       int    `json:"id" form:"id"`
	Category string `json:"category" form:"category"`
}

type EventCat struct {
	Event
	Categories Categories
}

type JoinedEvent struct {
	Event
	EventId int
	UserId  int
}
