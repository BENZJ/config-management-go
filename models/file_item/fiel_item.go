package fileitem

import (
	"config-management-go/models/file"
	"config-management-go/models/iteration"
	"config-management-go/utils"
)

type FieldItem struct {
	ID          int            `gorm:"primary_key;auto_increment" json:"id"`
	IterationID int            `gorm:"null" json:"iterationId"`
	FileID      int            `gorm:"null" json:"fileId"`
	Content     string         `gorm:"null" json:"content"`
	CreatedAt   utils.Datetime `gorm:"null" json:"createdAt"`
	CreatedBy   string         `gorm:"default:CURRENT_TIMESTAMP" json:"createdBy"`
	UpdatedAt   utils.Datetime `gorm:"null" json:"updatedAt"`
	UpdatedBy   string         `gorm:"null" json:"updatedBy"`
	Remark      string         `gorm:"null" json:"remark"`

	// Gorm 外键关联
	Iteration iteration.Iteration `gorm:"foreignkey:IterationID" json:"iteration"`
	File      file.File           `gorm:"foreignkey:FileID" json:"file"`
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
