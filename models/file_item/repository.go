package fileitem

import "github.com/jinzhu/gorm"

type repositoryImpl struct {
	db *gorm.DB
}

// Create implements Repository.
func (*repositoryImpl) Create(item *FieldItem) error {
	panic("unimplemented")
}

// DeleteItem implements Repository.
func (*repositoryImpl) DeleteItem(id int) error {
	panic("unimplemented")
}

// ListAll implements Repository.
func (*repositoryImpl) ListAll(items *[]FieldItem) error {
	panic("unimplemented")
}

// ModifyItem implements Repository.
func (*repositoryImpl) ModifyItem(item *FieldItem) error {
	panic("unimplemented")
}

func NewRepository(db *gorm.DB) Repository {
	return &repositoryImpl{
		db: db,
	}
}
