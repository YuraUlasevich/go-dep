package repo

import (
	"context"
	"database/sql"

	"github.com/go-dep/models"
)

type Repo struct {
	Conn *sql.DB
}

type UserRepository interface {
	GetUser(ctx context.Context, name string) (models.User, error)
	AddUser(ctx context.Context, name string) error
	DeleteUser(ctx context.Context, name string) error
	UpdateUser(ctx context.Context, oldName string, newName string) error
}

func (repo Repo) GetUser(ctx context.Context, name string) (models.User, error) {
	var user models.User

	rows, err := repo.Conn.Query("select * from users where name = $1", name)
	if err != nil {
		return models.User{}, err
	}

	defer rows.Close()

	for rows.Next() {

		err := rows.Scan(&user.ID, &user.Name)
		if err != nil {
			return models.User{}, err
		}
	}

	return user, nil
}

func (repo Repo) AddUser(ctx context.Context, name string) error {

	_, err := repo.Conn.ExecContext(ctx, "insert into users(name) values ($1)", name)
	if err != nil {
		return err
	}

	return nil

}

func (repo Repo) DeleteUser(ctx context.Context, name string) error {
	_, err := repo.Conn.ExecContext(ctx, "delete from users where name=$1", name)

	if err != nil {
		return err
	}

	return nil
}

func (repo Repo) UpdateUser(ctx context.Context, oldName string, newName string) error {
	_, err := repo.Conn.ExecContext(ctx, "update users set name=$1 where name=$2", newName, oldName)

	if err != nil {
		return err
	}

	return nil
}
