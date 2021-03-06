package database

import (
	// "fmt"
	"gorm.io/gorm"
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
	if err = repo.Select("Content").Save(looseLeaf).Error; err != nil {
		return
	}
	return
}

func (repo *LooseLeafRepository) UpdateBinderID(looseLeaf domain.LooseLeaf) (err error) {
	if looseLeaf.BinderID == (*uint)(nil) {
		err = repo.Model(&looseLeaf).Update("BinderID", gorm.Expr("NULL")).Error
	} else {
		err = repo.Model(&looseLeaf).Update("BinderID", looseLeaf.BinderID).Error
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

func (repo *LooseLeafRepository) FindAll(user domain.User, binderId *uint, page int, pageSize int) (looseLeafs []domain.LooseLeaf, err error) {
	var query *gorm.DB
	if binderId == (*uint)(nil) {
		query = repo.Where("user_id = ?", user.ID)
	} else {
		query = repo.Where("user_id = ?", user.ID).Where("binder_id = ?", binderId)
	}
	err = query.Order("ID desc").Offset((page - 1) * pageSize).Limit(pageSize).Find(&looseLeafs).Error
	return
}
