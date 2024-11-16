package user

import (
	"database/sql"
	"time"

	model "github.com/Jefschlarski/golang-modular/modules/user/model"
)

type UserRepository struct {
	connection *sql.DB
}

func NewUserRepository(connection *sql.DB) UserRepositoryInterface {
	return &UserRepository{connection: connection}
}

func (ur *UserRepository) GetUser(id int) (user *model.User, err error) {

	user = &model.User{}

	query, err := ur.connection.Prepare(`SELECT id, name, email, phone, user_type, created_at, updated_at FROM "user" WHERE id = $1`)
	if err != nil {
		return nil, err
	}

	defer query.Close()

	err = query.QueryRow(id).Scan(&user.ID, &user.Name, &user.Email, &user.Phone, &user.UserType, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return user, nil
}

func (ur *UserRepository) GetUsers() (users []*model.User, err error) {

	users = []*model.User{}

	query, err := ur.connection.Prepare(`SELECT id, name, email, phone, user_type, created_at, updated_at FROM "user"`)
	if err != nil {
		return nil, err
	}

	defer query.Close()

	rows, err := query.Query()
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		user := &model.User{}
		err = rows.Scan(&user.ID, &user.Name, &user.Email, &user.Phone, &user.UserType, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (ur *UserRepository) CreateUser(createUser model.User) (userId int, err error) {

	query, err := ur.connection.Prepare(`INSERT INTO "user" (name, email, phone, user_type, password, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`)
	if err != nil {
		return
	}

	defer query.Close()

	err = query.QueryRow(createUser.Name, createUser.Email, createUser.Phone, createUser.UserType, createUser.Password, createUser.CreatedAt, createUser.UpdatedAt).Scan(&userId)
	if err != nil {
		return
	}

	defer query.Close()

	return
}

func (ur *UserRepository) UpdateUser(id int, user model.User) (rowsAffected int64, err error) {

	query, err := ur.connection.Prepare(`UPDATE "user" SET name = $1, email = $2, phone = $3, user_type = $4, updated_at = $5 WHERE id = $6`)
	if err != nil {
		return
	}

	defer query.Close()

	result, err := query.Exec(user.Name, user.Email, user.Phone, user.UserType, user.UpdatedAt, id)
	if err != nil {
		return
	}

	rowsAffected, err = result.RowsAffected()
	if err != nil {
		return
	}

	return
}

func (ur *UserRepository) DeleteUser(id int) (rowsAffected int64, err error) {

	query, err := ur.connection.Prepare(`DELETE FROM "user" WHERE id = $1`)
	if err != nil {
		return
	}

	defer query.Close()

	result, err := query.Exec(id)
	if err != nil {
		return
	}

	rowsAffected, err = result.RowsAffected()
	if err != nil {
		return
	}

	return
}

func (ur *UserRepository) UpdatePassword(id int, newPassword string) (rowsAffected int64, err error) {

	query, err := ur.connection.Prepare(`UPDATE "user" SET password = $1, updated_at = $2 WHERE id = $3`)
	if err != nil {
		return
	}

	defer query.Close()

	result, err := query.Exec(newPassword, time.Now().Format("2006-01-02 15:04:05"), id)
	if err != nil {
		return
	}

	rowsAffected, err = result.RowsAffected()
	if err != nil {
		return
	}

	return
}

func (ur *UserRepository) GetUserPassword(id int) (password string, err error) {

	query, err := ur.connection.Prepare(`SELECT password FROM "user" WHERE id = $1`)
	if err != nil {
		return
	}

	defer query.Close()

	err = query.QueryRow(id).Scan(&password)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", nil
		}

		return "", err
	}

	return
}

func (ur *UserRepository) GetUserByEmail(email string) (user *model.User, err error) {

	user = &model.User{}

	query, err := ur.connection.Prepare(`SELECT id, name, email, phone, user_type, password FROM "user" WHERE email = $1`)
	if err != nil {
		return nil, err
	}

	defer query.Close()

	err = query.QueryRow(email).Scan(&user.ID, &user.Name, &user.Email, &user.Phone, &user.UserType, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return
}
