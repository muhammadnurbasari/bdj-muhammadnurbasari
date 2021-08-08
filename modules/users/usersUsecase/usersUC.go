package usersUsecase

import (
	"bdj-muhammadnurbasari/models/usersModel"
	"bdj-muhammadnurbasari/modules/users"
)

type UsersUsecase struct {
	usersrepo users.UsersRepository
}

func NewUsersUsecase(usersrepo users.UsersRepository) users.UsersUsecase {
	return &UsersUsecase{usersrepo: usersrepo}
}

func (usersUC *UsersUsecase) GetAllUsers() (*[]usersModel.Users, error) {
	resp, err := usersUC.usersrepo.RetrieveAllUsers()

	if err != nil {
		return nil, err
	}

	return resp, nil
}
