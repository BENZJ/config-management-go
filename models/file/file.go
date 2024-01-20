package file

import (
	"config-management-go/models/iteration"
	"time"
)

type File struct {
	ID          int       `gorm:"primary_key;auto_increment" json:"id"`
	IterationID int       `gorm:"null" json:"iterationId"`
	FileName    string    `gorm:"null" json:"fileName"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"createdAt"`
	Remark      string    `gorm:"null" json:"remark"`
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
	ListById(id int, file *File) error
}
