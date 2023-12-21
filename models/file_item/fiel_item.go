package fileitem

import (
	"config-management-go/models/file"
	"config-management-go/models/iteration"
	"time"
)

type FieldItem struct {
	ID          int       `gorm:"primary_key;auto_increment"`
	IterationID int       `gorm:"not null"`
	FileID      int       `gorm:"not null"`
	Content     string    `gorm:"not null"`
	CreatedAt   time.Time `gorm:"not null"`
	CreatedBy   string    `gorm:"not null"`
	UpdatedAt   time.Time `gorm:"null"`
	UpdatedBy   string    `gorm:"null"`

	// Gorm 外键关联
	Iteration iteration.Iteration `gorm:"foreignkey:IterationID"`
	File      file.File           `gorm:"foreignkey:FileID"`
}

func (fileItem *FieldItem) TableName() string {
	return "file_items"
}

type Repository interface {
	Create(item *FieldItem) error
	ListAll(fileId int, items *[]FieldItem) error
	DeleteItem(id int) error
	ModifyItem(item *FieldItem) error
}
