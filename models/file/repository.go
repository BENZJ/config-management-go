package file

import "github.com/jinzhu/gorm"

type repositoryImpl struct {
	db *gorm.DB
}

// Create implements Repository.
func (*repositoryImpl) Create(item *File) error {
	panic("unimplemented")
}

// DeleteItem implements Repository.
func (*repositoryImpl) DeleteItem(id int) error {
	panic("unimplemented")
}

// ListAll implements Repository.
func (*repositoryImpl) ListAll(files *[]File) error {
	panic("unimplemented")
}

// ModifyItem implements Repository.
func (*repositoryImpl) ModifyItem(file *File) error {
	panic("unimplemented")
}

func NewRepository(db *gorm.DB) Repository {
	return &repositoryImpl{
		db: db,
	}
}
