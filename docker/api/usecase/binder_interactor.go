package usecase

import (
	// "errors"
	"fmt"
	"myapp/domain"
)

type BinderInteractor struct {
	BinderRepository BinderRepository
}

type CustomError struct {
	Msg  string
	Code int
	error
}

func (err *CustomError) Error() string {
	return fmt.Sprintf("err %s [code=%d]", err.Msg, err.Code)
}

func (interactor *BinderInteractor) Add(binder *domain.Binder) (err error) {
	existingBinders, _ := interactor.Binders(binder.User, 1, 30)
	if len(existingBinders) >= 30 {
		err = &CustomError{Msg: "too many binders", Code: 400}
		return
	}
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
