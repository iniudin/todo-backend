package web

type TodoUpdateResponse struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	IsDone      bool   `json:"status"`
}
