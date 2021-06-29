package database

import "myapp/domain"

type UserRepository struct {
	SqlHandler
}

func (repo *UserRepository) Store(user domain.User) (id uint, err error) {
	if err = repo.Select("Uid", "Name").Create(&user).Error; err != nil {
		return
	}
	id = user.ID
	return
}

func (repo *UserRepository) FindByUid(uid string) (user domain.User, err error) {
	if err = repo.First(&user, "uid = ?", uid).Error; err != nil {
		return
	}
	return
}

func (repo *UserRepository) FindAll() (users []domain.User, err error) {
	if err = repo.Order("ID asc").Find(&users).Error; err != nil {
		return
	}
	return
}
