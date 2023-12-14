package request

import "time"

type Iteration struct {
	Name      string    `json:"name" binding:"required"`
	ID        int       `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
}
