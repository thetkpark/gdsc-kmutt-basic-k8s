package todo

type Todo struct {
	ID uint
	Name     string `json:"name"`
	Finished bool   `json:"finished"`
}
