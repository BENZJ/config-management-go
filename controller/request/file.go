package request

import "time"

type File struct {
	ID          int       `json:"id"`
	IterationID int       `json:"iterationId" binding:"required"`
	FileName    string    `json:"fileName" binding:"required"`
	CreatedAt   time.Time `json:"createdAt"`
}
