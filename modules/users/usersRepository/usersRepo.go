package usersRepository

import (
	"bdj-muhammadnurbasari/models/usersModel"
	"bdj-muhammadnurbasari/modules/users"
	"database/sql"
	"errors"
)

type sqlRepository struct {
	Conn *sql.DB
}

func NewUsersRepository(Conn *sql.DB) users.UsersRepository {
	return &sqlRepository{Conn: Conn}
}

func (db *sqlRepository) RetrieveAllUsers() (*[]usersModel.Users, error) {
	query := `SELECT id, name, position, job_desk, created_at FROM users WHERE deleted_at IS NULL`

	rows, errQuery := db.Conn.Query(query)
	if errQuery != nil {
		return nil, errors.New("usersRepository.RetrieveAllUsers errQuery = " + errQuery.Error())
	}

	defer rows.Close()

	var list []usersModel.Users

	for rows.Next() {
		var data usersModel.Users

		errScan := rows.Scan(&data.ID, &data.Name, &data.Position, &data.JobDesk, &data.CreatedAt)
		if errScan != nil {
			return nil, errors.New("usersRepository.RetrieveAllUsers errScan = " + errScan.Error())
		}

		list = append(list, data)
	}

	return &list, nil
}
