package web

type TodoUpdateRequest struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"       validate:"required"`
	Description string `json:"description" validate:"required"`
	IsDone      bool   `json:"is_done"`
}
