package showtodo

type showTodoViewModel struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	IsCompleted bool   `json:"isCompleted"`
}
