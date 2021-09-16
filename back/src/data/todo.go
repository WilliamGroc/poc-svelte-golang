package todoData

import (
	todoEntity "poc-svelte-golang/back/src/entity"
)

var Todos map[string]*todoEntity.Todo = make(map[string]*todoEntity.Todo)