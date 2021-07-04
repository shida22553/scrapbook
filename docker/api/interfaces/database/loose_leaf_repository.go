package database

import (
	// "fmt"
	"myapp/domain"
)

type LooseLeafRepository struct {
	SqlHandler
}

func (repo *LooseLeafRepository) Create(looseLeaf *domain.LooseLeaf) (err error) {
	if err = repo.Select("Content", "UserID").Create(looseLeaf).Error; err != nil {
		return
	}
	return
}

func (repo *LooseLeafRepository) Update(looseLeaf *domain.LooseLeaf) (err error) {
	if err = repo.Save(looseLeaf).Error; err != nil {
		return
	}
	return
}

func (repo *LooseLeafRepository) Delete(looseLeaf *domain.LooseLeaf) (err error) {
	if err = repo.Where("id = ?", looseLeaf.ID).Delete(looseLeaf).Error; err != nil {
		return
	}
	return
}

func (repo *LooseLeafRepository) FindById(user domain.User, id uint) (looseLeaf domain.LooseLeaf, err error) {
	if err = repo.Where("user_id = ? and id = ?", user.ID, id).Find(&looseLeaf).Error; err != nil {
		return
	}
	return
}

func (repo *LooseLeafRepository) FindAll(user domain.User, page int, pageSize int) (looseLeafs []domain.LooseLeaf, err error) {
	if err = repo.Where("user_id = ?", user.ID).Order("ID desc").Offset(page).Limit(pageSize).Find(&looseLeafs).Error; err != nil {
		return
	}
	return
}
