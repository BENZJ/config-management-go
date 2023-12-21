package request

import "time"

type FileItem struct {
	ID          int       `json:"id"`
	IterationID int       `json:"iterationID"`
	FileID      int       `json:"fileID"`
	Content     string    `json:"content"`
	CreatedAt   time.Time `json:"createdAt"`
	CreatedBy   int       `json:"createdBy"`
	UpdatedAt   time.Time `json:"updatedAt"`
	UpdatedBy   int       `json:"updatedBy"`
}
