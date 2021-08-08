package users

import "bdj-muhammadnurbasari/models/usersModel"

type UsersUsecase interface {
	GetAllUsers() (*[]usersModel.Users, error)
}
