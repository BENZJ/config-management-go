package file

import (
	"config-management-go/models/iteration"
	"time"
)

type File struct {
	ID          int       `gorm:"primary_key;auto_increment"`
	IterationID int       `gorm:"not null"`
	FileName    string    `gorm:"not null"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP"`

	// Gorm 外键关联
	Iteration iteration.Iteration `gorm:"foreignkey:IterationID"`
}

func (file *File) TableName() string {
	return "files"
}

type Repository interface {
	Create(item *File) error
	ListAll(iterationId int, files *[]File) error
	DeleteItem(id int) error
	ModifyItem(file *File) error
}
