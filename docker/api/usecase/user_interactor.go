package usecase

import "myapp/domain"

type UserInteractor struct {
	UserRepository UserRepository
}

func (interactor *UserInteractor) Add(user domain.User) (id uint, err error) {
	id, err = interactor.UserRepository.Create(user)
	return
}

func (interactor *UserInteractor) Users() (users []domain.User, err error) {
	users, err = interactor.UserRepository.FindAll()
	return
}

func (interactor *UserInteractor) UserByUid(uid string) (user domain.User, err error) {
	user, err = interactor.UserRepository.FindByUid(uid)
	return
}
