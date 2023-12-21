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
	return rep.db.Where("file_id= ?", fileId).Find(&items).Error
}

// ModifyItem implements Repository.
func (rep *repositoryImpl) ModifyItem(item *FieldItem) error {
	result := rep.db.Model(&FieldItem{}).Where("id = ?", item.ID).
		Updates(map[string]interface{}{
			"Content":   item.Content,
			"UpdatedBy": item.UpdatedBy,
		})
	return result.Error
}

func NewRepository(db *gorm.DB) Repository {
	return &repositoryImpl{
		db: db,
	}
}
