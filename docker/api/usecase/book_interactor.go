package usecase

import (
	"myapp/domain"
)

type BookInteractor struct {
	BookRepository BookRepository
}

func (interactor *BookInteractor) Add(book *domain.Book) (err error) {
	err = interactor.BookRepository.Create(book)
	return
}

func (interactor *BookInteractor) Remove(book *domain.Book) (err error) {
	err = interactor.BookRepository.Delete(book)
	return
}

func (interactor *BookInteractor) Update(book *domain.Book) (err error) {
	err = interactor.BookRepository.Update(book)
	return
}

func (interactor *BookInteractor) Books(user domain.User, page int, pageSize int) (books []domain.Book, err error) {
	books, err = interactor.BookRepository.FindAll(user, page, pageSize)
	return
}

func (interactor *BookInteractor) BookById(user domain.User, id uint) (book domain.Book, err error) {
	book, err = interactor.BookRepository.FindById(user, id)
	return
}
