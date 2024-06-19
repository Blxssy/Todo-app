package backend

type Todo struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	State bool   `json:"state"`
}
