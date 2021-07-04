package usecase

import "myapp/domain"

type BookRepository interface {
	Create(*domain.Book) error
	Update(*domain.Book) error
	Delete(*domain.Book) error
	FindById(domain.User, uint) (domain.Book, error)
	FindAll(domain.User, int, int) ([]domain.Book, error)
}
