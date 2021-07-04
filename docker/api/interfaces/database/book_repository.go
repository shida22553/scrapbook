package database

import (
	// "fmt"
	"myapp/domain"
)

type BookRepository struct {
	SqlHandler
}

func (repo *BookRepository) Create(book *domain.Book) (err error) {
	if err = repo.Select("Name", "UserID").Create(book).Error; err != nil {
		return
	}
	return
}

func (repo *BookRepository) Update(book *domain.Book) (err error) {
	if err = repo.Save(book).Error; err != nil {
		return
	}
	return
}

func (repo *BookRepository) Delete(book *domain.Book) (err error) {
	if err = repo.Where("id = ?", book.ID).Delete(book).Error; err != nil {
		return
	}
	return
}

func (repo *BookRepository) FindById(user domain.User, id uint) (book domain.Book, err error) {
	if err = repo.Where("user_id = ? and id = ?", user.ID, id).Find(&book).Error; err != nil {
		return
	}
	return
}

func (repo *BookRepository) FindAll(user domain.User, page int, pageSize int) (books []domain.Book, err error) {
	if err = repo.Where("user_id = ?", user.ID).Order("ID desc").Offset(page).Limit(pageSize).Find(&books).Error; err != nil {
		return
	}
	return
}
