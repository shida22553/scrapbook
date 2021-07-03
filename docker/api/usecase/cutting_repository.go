package usecase

import "myapp/domain"

type CuttingRepository interface {
	Create(*domain.Cutting) error
	Update(*domain.Cutting) error
	Delete(*domain.Cutting) error
	FindById(domain.User, uint) (domain.Cutting, error)
	FindAll(domain.User, int, int) ([]domain.Cutting, error)
}
