package usecase

import (
	"myapp/domain"
)

type BinderInteractor struct {
	BinderRepository BinderRepository
}

func (interactor *BinderInteractor) Add(binder *domain.Binder) (err error) {
	err = interactor.BinderRepository.Create(binder)
	return
}

func (interactor *BinderInteractor) Remove(binder *domain.Binder) (err error) {
	err = interactor.BinderRepository.Delete(binder)
	return
}

func (interactor *BinderInteractor) Update(binder *domain.Binder) (err error) {
	err = interactor.BinderRepository.Update(binder)
	return
}

func (interactor *BinderInteractor) Binders(user domain.User, page int, pageSize int) (binders []domain.Binder, err error) {
	binders, err = interactor.BinderRepository.FindAll(user, page, pageSize)
	return
}

func (interactor *BinderInteractor) BinderById(user domain.User, id uint) (binder domain.Binder, err error) {
	binder, err = interactor.BinderRepository.FindById(user, id)
	return
}
