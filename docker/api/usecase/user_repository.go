package usecase

import "myapp/domain"

type UserRepository interface {
	Store(domain.User) (uint, error)
	FindByUid(string) (domain.User, error)
	FindAll() ([]domain.User, error)
}
