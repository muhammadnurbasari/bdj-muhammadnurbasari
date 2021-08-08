package users

import "bdj-muhammadnurbasari/models/usersModel"

type UsersRepository interface {
	RetrieveAllUsers() (*[]usersModel.Users, error)
}
