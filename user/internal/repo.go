package internal

import (
	"errors"

	pg "github.com/go-pg/pg/v10"
	"github.com/google/uuid"

	"github.com/ssibrahimbas/go-uuid/entity"
)

type Repository struct {
	db *pg.DB
}

func NewRepository(db *pg.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) NewUser(u *entity.User) error {
	_, err := r.db.Model(u).Insert()
	return err
}

func (r *Repository) GetUserById(id uuid.UUID) (*entity.User, error) {
	usr := &entity.User{}
	err := r.db.Model(usr).Where("id = ?", id).Select()
	if err != nil {
		return nil, err
	}
	return usr, nil
}

func (r *Repository) UserIsExist(m string) (bool, error) {
	usr := &entity.User{
		Email: "",
	}
	err := r.db.Model(usr).Where("email = ?", m).Select()
	if err != nil {
		if err == pg.ErrNoRows {
			return false, nil
		}
		return false, err
	}
	return usr.Email != "", errors.New("User already exist")
}

func (r *Repository) GetAllUsers() ([]entity.User, error) {
	var arr []entity.User
	err := r.db.Model(&arr).Select()
	return arr, err
}
