package web

type TodoCreateRequest struct {
	Title       string `json:"title"       validate:"required"`
	Description string `json:"description" validate:"required"`
	IsDone      bool   `json:"is_done"`
}
