package client

type User struct {
	ID    int    `json:"id"`
	First string `json:"first_name"`
	Last  string `json:"second_name"`
	Email string `json:"email"`
}

func NewUser(id int, first string, last string, email string) *User {

	return &User{
		ID:    id,
		First: first,
		Last:  last,
		Email: email,
	}
}
