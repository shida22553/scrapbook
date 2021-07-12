package database

import (
	// "fmt"
	"myapp/domain"
)

type BinderRepository struct {
	SqlHandler
}

func (repo *BinderRepository) Create(binder *domain.Binder) (err error) {
	if err = repo.Select("Name", "UserID").Create(binder).Error; err != nil {
		return
	}
	return
}

func (repo *BinderRepository) Update(binder *domain.Binder) (err error) {
	if err = repo.Save(binder).Error; err != nil {
		return
	}
	return
}

func (repo *BinderRepository) Delete(binder *domain.Binder) (err error) {
	if err = repo.Where("id = ?", binder.ID).Delete(binder).Error; err != nil {
		return
	}
	return
}

func (repo *BinderRepository) FindById(user domain.User, id uint) (binder domain.Binder, err error) {
	if err = repo.Where("user_id = ? and id = ?", user.ID, id).Find(&binder).Error; err != nil {
		return
	}
	return
}

func (repo *BinderRepository) FindAll(user domain.User, page int, pageSize int) (binders []domain.Binder, err error) {
	if err = repo.Where("user_id = ?", user.ID).Order("ID desc").Offset(page - 1).Limit(pageSize).Find(&binders).Error; err != nil {
		return
	}
	return
}
