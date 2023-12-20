package iteration

import "github.com/jinzhu/gorm"

type repositoryImpl struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repositoryImpl{
		db: db,
	}
}

func (rep *repositoryImpl) Create(item *Iteration) error {
	return rep.db.Create(item).Error
}

func (rep *repositoryImpl) ListAll(iterations *[]Iteration) error {
	return rep.db.Find(iterations).Error
}
