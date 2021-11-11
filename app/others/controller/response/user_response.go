package others

import "github.com/fyk7/go-clean-arch-demo-v3/app/domain/model"

type CreateUserResponse struct {
	ID    string
	Email string
}

func ToCreateUserResp(usr *model.User) CreateUserResponse {
	return CreateUserResponse{
		ID:    usr.ID,
		Email: usr.Email,
	}
}
