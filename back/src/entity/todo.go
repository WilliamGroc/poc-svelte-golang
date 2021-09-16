package todoEntity

import (
	"github.com/google/uuid"
)

var orderCount = 0

type Todo struct {
	Id      string  `json:"id"`
	Text    string `json:"text"`
	Checked bool   `json:"checked"`
	Order int `json:"order"`
}

func MakeTodo() *Todo {
	orderCount++
	return &Todo{Id: uuid.NewString(), Order: orderCount}
}

func (t *Todo) Check() {
	t.Checked = true
}
