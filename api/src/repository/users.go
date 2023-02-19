package repository

import (
	"api/src/models"
	"database/sql"
)

type users struct {
	db *sql.DB
}

func NewUsersRepository(db *sql.DB) *users {
	return &users{db}
}

func (repository users) Create(user models.User) (uint64, error) {
	statement, err := repository.db.Prepare("insert into users (name, nickname, email, password) values (?, ?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(user.Name, user.Nickname, user.Email, user.Password)
	if err != nil {
		return 0, err
	}

	lastID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastID), nil
}

func (repository users) SearchByNameOrNick(nameOrNick string) ([]models.User, error) {
	nameOrNick = "%" + nameOrNick + "%"
	rows, err := repository.db.Query("select id, name, nickname, email, createAt from users where name like ? or nickname like ?", nameOrNick, nameOrNick)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		if err = rows.Scan(&user.ID, &user.Name, &user.Nickname, &user.Email, &user.CreatedAt); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (repository users) SearchByID(id uint64) (models.User, error) {
	row, err := repository.db.Query("select id, name, nickname, email, createAt from users where id = ?", id)

	if err != nil {
		return models.User{}, err
	}

	var user models.User

	if row.Next() {
		if err = row.Scan(&user.ID, &user.Name, &user.Nickname, &user.Email, &user.CreatedAt); err != nil {
			return models.User{}, err
		}
	}

	return user, nil
}

func (repository users) Update(id uint64, user models.User) error {
	statement, err := repository.db.Prepare("update users set name = ?, nickname = ?, email = ? where id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(user.Name, user.Nickname, user.Email, id); err != nil {
		return err
	}

	return nil
}
