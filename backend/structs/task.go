package requests

import "time"

type TaskRequest struct {
	Title       string    `json:"title" binding:"required"`
	Description string    `json:"description"`
	Status      string    `json:"status" binding:"required,oneof=todo inprogress done"`
	DueDate     time.Time `json:"due_date" binding:"required"`
}
