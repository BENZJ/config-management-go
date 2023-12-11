package iteration

import (
	"time"
)

type Iteration struct {
	ID        int       `gorm:"primary_key;AUTO_INCREMENT"`
	Name      string    `gorm:"not null;size:255"`
	CreatedAt time.Time `gorm:"not null"`
}

func (iteration *Iteration) TableName() string {
	return "iterations"
}

type Repository interface {
	Create(item *Iteration) error
}
