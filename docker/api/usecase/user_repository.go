package usecase

import "myapp/domain"

type UserRepository interface {
	Create(domain.User) (uint, error)
	Update(*domain.User) error
	FindByUid(string) (domain.User, error)
	FindAll() ([]domain.User, error)
}
