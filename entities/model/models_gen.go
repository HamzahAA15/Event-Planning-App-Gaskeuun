// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Comment struct {
	ID        int    `json:"id"`
	Comment   string `json:"comment"`
	User      *User  `json:"user"`
	UpdatedAt string `json:"updatedAt"`
}

type CommentsResponse struct {
	Comments  []*Comment `json:"comments"`
	TotalPage int        `json:"totalPage"`
}

type EditEvent struct {
	CategoryID  *int    `json:"categoryId"`
	Title       *string `json:"title"`
	Host        *string `json:"host"`
	Date        *string `json:"date"`
	Location    *string `json:"location"`
	Description *string `json:"description"`
	ImageURL    *string `json:"imageUrl"`
}

type EditUser struct {
	Name     *string `json:"name"`
	Email    *string `json:"email"`
	Password *string `json:"password"`
	ImageURL *string `json:"imageUrl"`
}

type Event struct {
	ID          *int    `json:"id"`
	UserID      int     `json:"userId"`
	CategoryID  int     `json:"categoryId"`
	Title       string  `json:"title"`
	Host        string  `json:"host"`
	Date        string  `json:"date"`
	Location    string  `json:"location"`
	Description string  `json:"description"`
	ImageURL    *string `json:"imageUrl"`
}

type LoginResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Token   string `json:"token"`
	User    *User  `json:"user"`
}

type NewEvent struct {
	UserID      int     `json:"userId"`
	CategoryID  int     `json:"categoryId"`
	Title       string  `json:"title"`
	Host        string  `json:"host"`
	Date        string  `json:"date"`
	Location    string  `json:"location"`
	Description string  `json:"description"`
	ImageURL    *string `json:"imageUrl"`
}

type NewUser struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ParticipantsResponse struct {
	Participants []*User `json:"participants"`
	TotalPage    int     `json:"totalPage"`
}

type SuccessResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type User struct {
	ID       *int    `json:"id"`
	Name     string  `json:"name"`
	Email    string  `json:"email"`
	Password string  `json:"password"`
	ImageURL *string `json:"imageUrl"`
}
