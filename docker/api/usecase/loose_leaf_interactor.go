package usecase

import (
	"myapp/domain"
)

type LooseLeafInteractor struct {
	LooseLeafRepository LooseLeafRepository
}

func (interactor *LooseLeafInteractor) Add(looseLeaf *domain.LooseLeaf) (err error) {
	err = interactor.LooseLeafRepository.Create(looseLeaf)
	return
}

func (interactor *LooseLeafInteractor) Remove(looseLeaf *domain.LooseLeaf) (err error) {
	err = interactor.LooseLeafRepository.Delete(looseLeaf)
	return
}

func (interactor *LooseLeafInteractor) Update(looseLeaf *domain.LooseLeaf) (err error) {
	err = interactor.LooseLeafRepository.Update(looseLeaf)
	return
}

func (interactor *LooseLeafInteractor) LooseLeafs(user domain.User, page int, pageSize int) (looseLeafs []domain.LooseLeaf, err error) {
	looseLeafs, err = interactor.LooseLeafRepository.FindAll(user, page, pageSize)
	return
}

func (interactor *LooseLeafInteractor) LooseLeafById(user domain.User, id uint) (looseLeaf domain.LooseLeaf, err error) {
	looseLeaf, err = interactor.LooseLeafRepository.FindById(user, id)
	return
}
