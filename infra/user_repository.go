package infra

import (
	"database/sql"

	"github.com/shota-tech/ddd-demo/domain/model"
	"github.com/shota-tech/ddd-demo/domain/repository"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) repository.UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) FindById(id int) (*model.User, error) {
	user := model.User{}
	sql := "SELECT id, name, email FROM users WHERE id = $1"
	err := r.db.QueryRow(sql, id).Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) FindAll() ([]model.User, error) {
	users := []model.User{}
	sql := "SELECT id, name, email FROM users"
	rows, err := r.db.Query(sql)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		user := model.User{}
		err = rows.Scan(&user.ID, &user.Name, &user.Email)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (r *userRepository) Create(user *model.User) (int, error) {
	sql := "INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id"
	err := r.db.QueryRow(sql, user.Name, user.Email).Scan(&user.ID)
	if err != nil {
		return 0, err
	}
	return user.ID, nil
}
