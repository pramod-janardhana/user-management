package domain

type User struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserStore interface {
	Create(user *User) error
	Get(Id string) (*User, error)
	// Update(Id string, user *User) (*User, error)
	// Delete(Id string) error
}
