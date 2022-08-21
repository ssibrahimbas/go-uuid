package internal

import (
	"errors"

	"github.com/google/uuid"
	"github.com/ssibrahimbas/go-uuid/entity"
)

type Service struct {
	repo *Repository
}

func NewService(r *Repository) *Service {
	return &Service{
		repo: r,
	}
}

func (s *Service) NewUser(dto *entity.CreateUserDto) (*entity.User, error) {
	if s.userIsExist(dto.Email) {
		return nil, errors.New("Email already exist")
	}
	usr := entity.CreateUser(dto.Email, dto.Password)
	err := s.repo.NewUser(usr)
	if err != nil {
		return nil, err
	}
	return usr, nil
}

func (s *Service) userIsExist(m string) bool {
	res, err := s.repo.UserIsExist(m)
	if err != nil {
		panic(err)
	}
	return res
}

func (s *Service) GetUserById(id string) (*entity.User, error) {
	i, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	return s.repo.GetUserById(i)
}

func (s *Service) GetAllUsers() ([]entity.User, error) {
	return s.repo.GetAllUsers()
}
