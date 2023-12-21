package file

import "github.com/jinzhu/gorm"

type repositoryImpl struct {
	db *gorm.DB
}

// Create implements Repository.
func (rep *repositoryImpl) Create(item *File) error {
	return rep.db.Create(item).Error
}

// DeleteItem implements Repository.
func (rep *repositoryImpl) DeleteItem(id int) error {
	panic("unimplemented")
}

// ListAll implements Repository.
func (rep *repositoryImpl) ListAll(iterationId int, files *[]File) error {
	return rep.db.Where("iteration_id= ?", iterationId).Find(&files).Error
}

// ModifyItem implements Repository.
func (rep *repositoryImpl) ModifyItem(file *File) error {
	panic("unimplemented")
}

func NewRepository(db *gorm.DB) Repository {
	return &repositoryImpl{
		db: db,
	}
}
