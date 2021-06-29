package database

import "myapp/domain"

type CuttingRepository struct {
	SqlHandler
}

func (repo *CuttingRepository) Create(cutting domain.Cutting) (id uint, err error) {
	if err = repo.Select("Note", "UserID").Create(&cutting).Error; err != nil {
		return
	}
	id = cutting.ID
	return
}

func (repo *CuttingRepository) Update(cutting domain.Cutting) (id uint, err error) {
	if err = repo.Save(&cutting).Error; err != nil {
		return
	}
	id = cutting.ID
	return
}

func (repo *CuttingRepository) FindById(user domain.User, id int) (cutting domain.Cutting, err error) {
	if err = repo.Where("user_id = ? and id = ?", user.ID, id).Find(&cutting).Error; err != nil {
		return
	}
	return
}

func (repo *CuttingRepository) FindAll(user domain.User, page int, pageSize int) (cuttings []domain.Cutting, err error) {
	if err = repo.Where("user_id = ?", user.ID).Order("ID desc").Offset(page).Limit(pageSize).Find(&cuttings).Error; err != nil {
		return
	}
	return
}
