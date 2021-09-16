package todoService

import (
	todoData "poc-svelte-golang/back/src/data"
	todoEntity "poc-svelte-golang/back/src/entity"
	"sort"
)


func Create(text string) {
	newTodo := todoEntity.MakeTodo()
	newTodo.Text = text

	Add(newTodo)
}

func Add(todo *todoEntity.Todo) {
	todoData.Todos[todo.Id] = todo
}

func Check(id string) *todoEntity.Todo {
	todoToUpdate := todoData.Todos[id]
	todoToUpdate.Check()

	return todoToUpdate
}

func GetMany() []*todoEntity.Todo {
		arr := make([]*todoEntity.Todo, 0)
	
		for _, v := range todoData.Todos {
			arr = append(arr, v)
		}
		
		sortByOrder(arr)
		return arr
}

type ByOrder []*todoEntity.Todo

func (a ByOrder) Len() int           { return len(a) }
func (a ByOrder) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByOrder) Less(i, j int) bool { return a[i].Order < a[j].Order }

func sortByOrder(arr []*todoEntity.Todo) {
	sort.Sort(ByOrder(arr))
}