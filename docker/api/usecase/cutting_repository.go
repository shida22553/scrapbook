package usecase

import "myapp/domain"

type CuttingRepository interface {
	Create(domain.Cutting) (uint, error)
	Update(domain.Cutting) (uint, error)
	FindById(domain.User, int) (domain.Cutting, error)
	FindAll(domain.User, int, int) ([]domain.Cutting, error)
}
