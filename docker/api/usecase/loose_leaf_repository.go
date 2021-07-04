package usecase

import "myapp/domain"

type LooseLeafRepository interface {
	Create(*domain.LooseLeaf) error
	Update(*domain.LooseLeaf) error
	UpdateBinderID(domain.LooseLeaf) error
	Delete(*domain.LooseLeaf) error
	FindById(domain.User, uint) (domain.LooseLeaf, error)
	FindAll(domain.User, int, int) ([]domain.LooseLeaf, error)
}
