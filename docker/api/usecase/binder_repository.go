package usecase

import "myapp/domain"

type BinderRepository interface {
	Create(*domain.Binder) error
	Update(*domain.Binder) error
	Delete(*domain.Binder) error
	FindById(domain.User, uint) (domain.Binder, error)
	FindAll(domain.User, int, int) ([]domain.Binder, error)
}
