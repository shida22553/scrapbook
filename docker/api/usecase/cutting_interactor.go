package usecase

import (
	"myapp/domain"
)

type CuttingInteractor struct {
	CuttingRepository CuttingRepository
}

func (interactor *CuttingInteractor) Add(cutting *domain.Cutting) (err error) {
	err = interactor.CuttingRepository.Create(cutting)
	return
}

func (interactor *CuttingInteractor) Remove(cutting *domain.Cutting) (err error) {
	err = interactor.CuttingRepository.Delete(cutting)
	return
}

func (interactor *CuttingInteractor) Update(cutting *domain.Cutting) (err error) {
	err = interactor.CuttingRepository.Update(cutting)
	return
}

func (interactor *CuttingInteractor) Cuttings(user domain.User, page int, pageSize int) (cuttings []domain.Cutting, err error) {
	cuttings, err = interactor.CuttingRepository.FindAll(user, page, pageSize)
	return
}

func (interactor *CuttingInteractor) CuttingById(user domain.User, id uint) (cutting domain.Cutting, err error) {
	cutting, err = interactor.CuttingRepository.FindById(user, id)
	return
}
