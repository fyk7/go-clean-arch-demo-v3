package model

type User struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}

// func NewUser(id, email string) *User {
// 	return &User{
// 		ID:    id,
// 		Email: email,
// 	}
// }

// func (u *User) GetID() string {
// 	return u.id
// }

// func (u *User) GetEmail() string {
// 	return u.email
// }
