package iteration

import (
	"time"
)

type Iteration struct {
	ID        int       `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	Name      string    `gorm:"not null;size:255" json:"name"`
	CreatedAt time.Time `gorm:"null" json:"created_at"`
}

func (iteration *Iteration) TableName() string {
	return "iterations"
}

type Repository interface {
	Create(item *Iteration) error
	ListAll(iterations *[]Iteration) error
}
