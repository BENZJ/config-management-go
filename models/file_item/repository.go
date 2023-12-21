package fileitem

import "github.com/jinzhu/gorm"

type repositoryImpl struct {
	db *gorm.DB
}

// Create implements Repository.
func (rep *repositoryImpl) Create(item *FieldItem) error {
	return rep.db.Create(item).Error
}

// DeleteItem implements Repository.
func (rep *repositoryImpl) DeleteItem(id int) error {
	panic("unimplemented")
}

// ListAll implements Repository.
func (rep *repositoryImpl) ListAll(fileId int, items *[]FieldItem) error {
	return rep.db.Where("iteration_id= ?", fileId).Find(&items).Error
}

// ModifyItem implements Repository.
func (rep *repositoryImpl) ModifyItem(item *FieldItem) error {
	return rep.db.Save(&item).Error
}

func NewRepository(db *gorm.DB) Repository {
	return &repositoryImpl{
		db: db,
	}
}
