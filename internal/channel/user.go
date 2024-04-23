package channel

type User struct {
	Id        int
	FirstName string
	LastName  string
	Email     string
}

func NewUser(id int, firstName string, lastName string, email string) *User {
	return &User{
		Id:        id,
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
	}
}
